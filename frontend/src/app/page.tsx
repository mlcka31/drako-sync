"use client";

import { ConnectBtn } from "~/components/ConnectButton";

import Profile from "~/components/Profile";
import SendMoneyButton from "~/components/SendMoneyButton";
import { useGetMessages } from "~/hooks/useGetMessages";
export default function HomePage() {
  const messages = useGetMessages();
  console.log(messages);
  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white">
      <div className="container flex flex-col items-center justify-center gap-12 px-4 py-16">
        <h1 className="text-5xl font-extrabold tracking-tight text-white sm:text-[5rem]">
          Create <span className="text-[hsl(280,100%,70%)]">T3</span> App
        </h1>
        <div className="z-50 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
          <ConnectBtn />
        </div>
        <SendMoneyButton />
        <Profile />
        <div>
          {messages?.data?.map((item) => {
            return (
              <div key={item.timestamp}>
                <p>{item.content}</p>
              </div>
            );
          })}
        </div>
      </div>
    </main>
  );
}
