"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.blockchainService = void 0;
const ethers_1 = require("ethers");
const AlephGameState_1 = require("../abis/AlephGameState");
const openaiService_1 = require("./openaiService");
const logger_1 = __importDefault(require("../middleware/logger"));
const walletService_1 = require("./walletService");
class BlockchainService {
    constructor(walletService) {
        this.isPolling = false;
        this.pollingInterval = null;
        // Initialize with default values, will be properly set in init()
        this.provider = new ethers_1.ethers.JsonRpcProvider(process.env.RPC_URL);
        this.walletService = walletService;
        this.wallet = new ethers_1.ethers.Wallet(walletService.getWallet().privateKey, this.provider);
        this.contractAddress = process.env.CONTRACT_ADDRESS;
        this.contract = new ethers_1.ethers.Contract(this.contractAddress, AlephGameState_1.AlephGameStateAbi, this.wallet);
        this.systemPrompt =
            process.env.SYSTEM_PROMPT ||
                "You are an AI assistant in a blockchain game.";
    }
    async init(contractAddress) {
        try {
            this.provider = new ethers_1.ethers.JsonRpcProvider(process.env.RPC_URL);
            this.wallet = new ethers_1.ethers.Wallet(this.walletService.getWallet().privateKey, this.provider);
            this.contract = new ethers_1.ethers.Contract(contractAddress, AlephGameState_1.AlephGameStateAbi, this.wallet);
            this.systemPrompt =
                process.env.SYSTEM_PROMPT ||
                    "You are an AI assistant in a blockchain game.";
            logger_1.default.info(`Blockchain service initialized with contract: ${contractAddress}`);
            // Verify we can connect to the contract
            const gameState = await this.contract.gameState();
            logger_1.default.info(`Current game state: ${gameState}`);
            return true;
        }
        catch (error) {
            logger_1.default.error("Failed to initialize blockchain service", error);
            return false;
        }
    }
    startPolling(intervalMs = 1000) {
        if (this.isPolling)
            return;
        this.isPolling = true;
        logger_1.default.info(`Starting blockchain polling with interval: ${intervalMs}ms`);
        this.pollingInterval = setInterval(async () => {
            try {
                await this.checkGameState();
            }
            catch (error) {
                logger_1.default.error("Error during blockchain polling", error);
            }
        }, intervalMs);
    }
    stopPolling() {
        if (this.pollingInterval) {
            clearInterval(this.pollingInterval);
            this.pollingInterval = null;
            this.isPolling = false;
            logger_1.default.info("Blockchain polling stopped");
        }
    }
    async checkGameState() {
        try {
            // Get current game state
            const gameState = await this.contract.gameState();
            // Check if it's agent's turn (assuming there's a specific state for this)
            // You may need to adjust this logic based on your contract's state enum
            if (gameState != 1n)
                return;
            // Assuming 1 is the state when it's agent's turn
            logger_1.default.info("It is agent's turn, generating response...");
            // Get messages to generate context for the AI
            const messages = await this.contract.getMessages();
            // Format messages for OpenAI
            const formattedMessages = await Promise.all([
                ...messages.map(async (msg) => ({
                    role: msg.sender.toLowerCase() ===
                        (await this.contract.aiAgentAddress()).toLowerCase()
                        ? "assistant"
                        : "user",
                    content: msg.content,
                })),
            ]);
            const response = await this.generateAIResponse(formattedMessages);
            if (response.isPayout) {
                await this.contract.payoutPrize();
            }
            if (response.content) {
                await this.sendReply(response.content);
            }
            else {
                logger_1.default.debug(`Not agent's turn. Current game state: ${gameState}`);
                throw new Error("No message or tool call found");
            }
        }
        catch (error) {
            logger_1.default.error("Error checking game state", error);
        }
    }
    async generateAIResponse(messages) {
        try {
            const lastMessages = messages.slice(-10);
            const updatedMessages = [
                { role: "system", content: this.systemPrompt },
                ...lastMessages,
            ];
            const completion = await openaiService_1.openai.chat.completions.create({
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
            logger_1.default.info(`Generated AI response: ${response}`);
            return response;
        }
        catch (error) {
            logger_1.default.error("Error generating AI response", error);
            throw error;
            // return "Sorry, I encountered an error while processing your request.";
        }
    }
    async sendReply(content) {
        try {
            logger_1.default.info("Sending reply to blockchain...");
            const tx = await this.contract.reply(content);
            const receipt = await tx.wait();
            logger_1.default.info(`Reply sent successfully. Transaction hash: ${receipt?.hash}`);
            return receipt;
        }
        catch (error) {
            logger_1.default.error("Error sending reply to blockchain", error);
            throw error;
        }
    }
}
exports.blockchainService = new BlockchainService(walletService_1.walletService);
