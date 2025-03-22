"use client";

import { utils } from "ethers";
import { ConnectBtn } from "~/components/ConnectButton";
import Logo from "~/components/Logo";
import Profile from "~/components/Profile";
import SendMoneyButton from "~/components/SendMoneyButton";
import { useGame } from "~/hooks/useGame";
import Link from "next/link";
import { useState } from "react";
import FigmaButton from "~/components/FigmaButton";
import Navbar from "~/components/Navbar";
import Chat from "~/components/Chat";
import { useAccount } from "wagmi";
import { useConnectModal } from "@rainbow-me/rainbowkit";
import { GameState } from "~/types/GameState";

export default function HomePage() {
  const [isStarted, setIsStarted] = useState(false);
  const {
    messages,
    prizePool,
    messagePrice,
    gameState,
    isLoading,
    isError,
    agentAddress,
    sendMessage,
  } = useGame();

  const { address } = useAccount();
  const { openConnectModal } = useConnectModal();
  console.log(messages);

  const handleSendMessage = (message: string) => {
    if (!address) {
      openConnectModal && openConnectModal();
      return;
    }
    if (!messagePrice) {
      // toast.error("Message price is not set");
      return;
    }
    sendMessage(message, utils.formatEther(messagePrice));
  };
  return (
    <main className="flex h-full w-full flex-col items-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white">
      {!isStarted ? (
        <div className="container z-10 mt-[100px] flex h-[90vh] flex-col items-center justify-center gap-12 px-4 py-16">
          <Logo className="h-20 w-full" />
          <div className="z-50 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
            <button
              className="h-[80px] w-full bg-transparent text-[32px] md:text-[80px]"
              onClick={() => {
                setIsStarted(true);
              }}
            >
              Start
            </button>
          </div>
        </div>
      ) : (
        <div className="container z-10 mt-[100px] flex h-full max-w-[800px] flex-grow flex-col gap-12">
          <Navbar />
          <Chat
            prizePool={prizePool?.toString() || "0"}
            messagePrice={messagePrice?.toString() || "0"}
            gameState={gameState as GameState}
            messages={
              messages?.map((item, index) => ({
                id: item.timestamp.toString(),
                content: item.content,
                role: item.sender === agentAddress.data ? "assistant" : "user",
                timestamp: new Date(Number(item.timestamp)),
              })) || []
            }
            onSendMessage={handleSendMessage}
            className="h-full w-full"
            isDisabled={isLoading || gameState !== GameState.UserAction}
          />
        </div>
      )}
    </main>
  );
}
