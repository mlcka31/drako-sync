import { useReadContract } from "wagmi";
import AlephGameState from "../abis/AlephGameState.json";
import { contractAddress } from "~/config";
import { Message } from "~/types/Message";
export const useGetMessages = () => {
  const result = useReadContract({
    abi: AlephGameState,
    address: contractAddress,
    functionName: "getMessages",
  });
  console.log(result);
  return {
    data: result.data as Message[],
    isLoading: result.isLoading,
    isError: result.isError,
  };
};
