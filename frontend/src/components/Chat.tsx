import React, { useState, useRef, useEffect } from "react";
import Image from "next/image";
import { GameState } from "~/types/GameState";
import { convertWei } from "~/utils/convertWei";

// Define the message type
export interface Message {
  id: string;
  content: string;
  role: "user" | "assistant";
  timestamp?: Date;
}

const formatGameState = (gameState: keyof typeof GameState) => {
  console.log("gameState", gameState);
  switch (gameState) {
    case "UserAction":
      return "User Turn";
    case "AgentAction":
      return "Agent Turn";
    case "Complete":
      return "Game Over";
    default:
      return "Uninitialized";
  }
};

interface ChatProps {
  messages: Message[];
  onSendMessage: (message: string) => void;
  className?: string;
  isDisabled?: boolean;
  prizePool: string;
  messagePrice: string;
  gameState: keyof typeof GameState;
}

const Chat: React.FC<ChatProps> = ({
  messages,
  onSendMessage,
  className,
  isDisabled,
  prizePool,
  messagePrice,
  gameState,
}) => {
  const [input, setInput] = useState("");
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (input.trim()) {
      onSendMessage(input);
      setInput("");
    }
  };

  return (
    <div
      className={`flex flex-col rounded-lg shadow-md ${className || "h-full"} overflow-hidden`}
    >
      <div className="flex-1 space-y-4 overflow-y-auto p-4">
        {messages.map((message) => (
          <div
            key={message.id}
            className={`flex ${message.role === "user" ? "justify-end" : "justify-start"}`}
          >
            <div
              className={`max-w-[80%] rounded-lg p-3 ${
                message.role === "user"
                  ? "rounded-tr-none bg-[#07071666] text-white"
                  : "rounded-tl-none bg-[#107378] text-white"
              }`}
            >
              <p>{message.content}</p>
              {message.timestamp && (
                <div
                  className={`mt-1 text-xs ${message.role === "user" ? "text-blue-100" : "text-gray-500"}`}
                >
                  {message.timestamp.toLocaleTimeString([], {
                    hour: "2-digit",
                    minute: "2-digit",
                  })}
                </div>
              )}
            </div>
          </div>
        ))}
        <div ref={messagesEndRef} />
      </div>

      <div
        className="p-4"
        style={{ backgroundColor: "#07071666", borderRadius: "16px" }}
      >
        <div className="flex w-full items-center justify-between gap-4 pb-2">
          <p className="text-[14px] text-white md:text-[24px]">
            Prize: {convertWei(prizePool)}
          </p>
          <p className="text-[14px] text-white md:text-[24px]">
            Message price: {convertWei(messagePrice)}
          </p>
          <p className="text-[14px] text-white md:text-[24px]">
            Game State: {formatGameState(gameState)}
          </p>
        </div>
        <form onSubmit={handleSubmit} className="flex items-center">
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Type a message..."
            className="flex-1 px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            style={{
              backgroundColor: "#F7F9FB0D",
              borderRadius: "16px",
            }}
            disabled={isDisabled}
          />
          <button
            type="submit"
            disabled={!input.trim() || isDisabled}
            className="ml-2 p-2 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50"
          >
            <Image
              src="/assets/send-icon.svg"
              alt="Send"
              width={24}
              height={24}
            />
          </button>
        </form>
      </div>
    </div>
  );
};

export default Chat;
