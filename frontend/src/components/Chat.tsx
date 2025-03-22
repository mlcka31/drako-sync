import React from "react";
import Image from "next/image";

// Define the message type
export interface Message {
  id: string;
  content: string;
  role: "user" | "assistant";
  timestamp?: Date;
}

interface ChatProps {
  messages: Message[];
  onSendMessage?: (message: string) => void;
  className?: string;
}

const Chat: React.FC<ChatProps> = ({ messages, onSendMessage, className }) => {
  const [inputValue, setInputValue] = React.useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (inputValue.trim() && onSendMessage) {
      onSendMessage(inputValue);
      setInputValue("");
    }
  };

  return (
    <div
      className={`flex flex-col rounded-lg shadow-md ${className || "h-full"} flex-grow`}
    >
      {/* Chat messages */}
      <div className="h-full w-full flex-1 flex-grow space-y-4 overflow-y-auto p-4">
        {messages.map((message) => (
          <div
            key={message.id}
            className={`flex ${message.role === "user" ? "justify-end" : "justify-start"}`}
          >
            <div
              className={`max-w-[80%] rounded-lg p-3 ${
                message.role === "user"
                  ? "rounded-tr-none bg-[#0707160D] text-white"
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
      </div>

      {/* Input area - now with sticky positioning */}
      <div
        className="sticky bottom-0 p-4"
        style={{ backgroundColor: "#07071666", borderRadius: "16px" }}
      >
        <form onSubmit={handleSubmit} className="flex items-center">
          <input
            type="text"
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            placeholder="Type a message..."
            className="flex-1 px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            style={{
              backgroundColor: "#F7F9FB0D",
              borderRadius: "16px",
            }}
          />
          <button
            type="submit"
            disabled={!inputValue.trim()}
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
