import { ethers } from "ethers";
import { AlephGameStateAbi } from "../abis/AlephGameState";
import { openai } from "./openaiService";
import logger from "../middleware/logger";
import { ChatCompletionMessageParam } from "openai/resources/chat";
import { AlephGameState } from "../typechain/AlephGameState";
import { walletService, WalletService } from "./walletService";
class BlockchainService {
  private provider: ethers.JsonRpcProvider;
  private contract: AlephGameState;
  private wallet: ethers.Wallet;
  private isPolling: boolean = false;
  private pollingInterval: NodeJS.Timeout | null = null;
  private systemPrompt: string;
  private contractAddress: string;
  private walletService: WalletService;

  constructor(walletService: WalletService) {
    // Initialize with default values, will be properly set in init()
    this.provider = new ethers.JsonRpcProvider(process.env.RPC_URL);
    this.walletService = walletService;
    this.wallet = new ethers.Wallet(
      walletService.getWallet().privateKey,
      this.provider
    );

    this.contractAddress = process.env.CONTRACT_ADDRESS as string;

    this.contract = new ethers.Contract(
      this.contractAddress,
      AlephGameStateAbi,
      this.wallet
    ) as unknown as AlephGameState;

    this.systemPrompt =
      process.env.SYSTEM_PROMPT ||
      "You are an AI assistant in a blockchain game.";
  }

  public async init(contractAddress: string) {
    try {
      this.provider = new ethers.JsonRpcProvider(process.env.RPC_URL);
      this.wallet = new ethers.Wallet(
        this.walletService.getWallet().privateKey,
        this.provider
      );
      this.contract = new ethers.Contract(
        contractAddress,
        AlephGameStateAbi,
        this.wallet
      ) as unknown as AlephGameState;
      this.systemPrompt =
        process.env.SYSTEM_PROMPT ||
        "You are an AI assistant in a blockchain game.";

      logger.info(
        `Blockchain service initialized with contract: ${contractAddress}`
      );

      // Verify we can connect to the contract
      const gameState = await this.contract.gameState();
      logger.info(`Current game state: ${gameState}`);

      return true;
    } catch (error) {
      logger.error("Failed to initialize blockchain service", error);
      return false;
    }
  }

  public startPolling(intervalMs: number = 1000) {
    if (this.isPolling) return;

    this.isPolling = true;
    logger.info(`Starting blockchain polling with interval: ${intervalMs}ms`);

    this.pollingInterval = setInterval(async () => {
      try {
        await this.checkGameState();
      } catch (error) {
        logger.error("Error during blockchain polling", error);
      }
    }, intervalMs);
  }

  public stopPolling() {
    if (this.pollingInterval) {
      clearInterval(this.pollingInterval);
      this.pollingInterval = null;
      this.isPolling = false;
      logger.info("Blockchain polling stopped");
    }
  }

  private async checkGameState() {
    try {
      // Get current game state
      const gameState = await this.contract.gameState();

      // Check if it's agent's turn (assuming there's a specific state for this)
      // You may need to adjust this logic based on your contract's state enum
      if (gameState != 1n) return;
      // Assuming 1 is the state when it's agent's turn
      logger.info("It is agent's turn, generating response...");

      // Get messages to generate context for the AI
      const messages = await this.contract.getMessages();
      // Format messages for OpenAI
      const formattedMessages = await Promise.all([
        ...messages.map(async (msg) => ({
          role:
            msg.sender.toLowerCase() ===
            (await this.contract.aiAgentAddress()).toLowerCase()
              ? ("assistant" as const)
              : ("user" as const),
          content: msg.content,
        })),
      ]);
      const response = await this.generateAIResponse(formattedMessages);

      if (response.isPayout) {
        await this.contract.payoutPrize();
      }
      if (response.content) {
        await this.sendReply(response.content);
      } else {
        logger.debug(`Not agent's turn. Current game state: ${gameState}`);
        throw new Error("No message or tool call found");
      }
    } catch (error) {
      logger.error("Error checking game state", error);
    }
  }

  private async generateAIResponse(
    messages: ChatCompletionMessageParam[]
  ): Promise<{
    isPayout: boolean;
    content?: string | null;
  }> {
    try {
      const lastMessages = messages.slice(-10);

      const updatedMessages: ChatCompletionMessageParam[] = [
        { role: "system", content: this.systemPrompt },
        ...lastMessages,
      ];

      const completion = await openai.chat.completions.create({
        model: process.env.OPENAI_MODEL || "gpt-4",
        messages: updatedMessages,
        max_tokens: 1000,
        tools: [
          {
            type: "function",
            function: {
              name: "payout_prize",
              description: "Payout the player the prize when they win",
            },
          },
        ],
      });

      const toolCall = completion.choices[0].message.tool_calls?.[0];
      const message = completion.choices[0].message.content;
      if (!message && !toolCall) {
        throw new Error("No message or tool call found");
      }
      const response = {
        isPayout: toolCall?.function.name === "payout_prize",
        content: message,
      };

      logger.info(`Generated AI response: ${response}`);

      return response;
    } catch (error) {
      logger.error("Error generating AI response", error);
      throw error;
      // return "Sorry, I encountered an error while processing your request.";
    }
  }

  private async sendReply(content: string) {
    try {
      logger.info("Sending reply to blockchain...");
      const tx = await this.contract.reply(content);
      const receipt = await tx.wait();
      logger.info(
        `Reply sent successfully. Transaction hash: ${receipt?.hash}`
      );
      return receipt;
    } catch (error) {
      logger.error("Error sending reply to blockchain", error);
      throw error;
    }
  }
}

export const blockchainService = new BlockchainService(walletService);
