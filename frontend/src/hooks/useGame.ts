import { useReadContract, useWriteContract } from "wagmi";
import { AlephGameStateAbi } from "~/abis/AlephGameState";
import { contractAddress } from "~/config";
import { GameState } from "~/types/GameState";
import { utils } from "ethers";

const pollingInterval = 1000;
export const useGame = () => {
  const messages = useReadContract({
    abi: AlephGameStateAbi,
    address: contractAddress,
    functionName: "getMessages",
    query: {
      refetchInterval: pollingInterval,
    },
  });

  const { writeContract } = useWriteContract();

  const sendMessage = (message: string, price: string) => {
    console.log(price);
    console.log(utils.parseEther(price).toString());
    writeContract({
      abi: AlephGameStateAbi,
      address: contractAddress,
      functionName: "sendMessage",
      value: BigInt(utils.parseEther(price).toString()),
      args: [message],
    });
  };

  const agentAddress = useReadContract({
    abi: AlephGameStateAbi,
    address: contractAddress,
    functionName: "aiAgentAddress",
    query: {
      refetchInterval: pollingInterval,
    },
  });

  const prizePool = useReadContract({
    abi: AlephGameStateAbi,
    address: contractAddress,
    functionName: "prizePool",
    query: {
      refetchInterval: pollingInterval,
    },
  });

  const gameState = useReadContract({
    abi: AlephGameStateAbi,
    address: contractAddress,
    functionName: "gameState",
    query: {
      refetchInterval: pollingInterval,
    },
  });

  const messagePrice = useReadContract({
    abi: AlephGameStateAbi,
    address: contractAddress,
    functionName: "messagePrice",
    query: {
      refetchInterval: pollingInterval,
    },
  });

  const isLoading =
    messages.isLoading || prizePool.isLoading || messagePrice.isLoading;
  const isError = messages.isError || prizePool.isError || messagePrice.isError;
  return {
    sendMessage,
    messages: messages.data,
    prizePool: prizePool.data,
    messagePrice: messagePrice.data,
    gameState: gameState.data && GameState[gameState.data],
    isLoading,
    isError,
    agentAddress: agentAddress.data,
  };
};
