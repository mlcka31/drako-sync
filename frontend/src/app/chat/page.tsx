"use client";

import React, { useState } from "react";
import Chat, { Message } from "~/components/Chat";
import { v4 as uuidv4 } from "uuid"; // You'll need to install this package

export default function ChatPage() {
  const [messages, setMessages] = useState<Message[]>([
    {
      id: "1",
      content: "Hello! How can I help you today?",
      role: "assistant",
      timestamp: new Date(),
    },
  ]);

  const handleSendMessage = (content: string) => {
    // Add user message
    const userMessage: Message = {
      id: uuidv4(),
      content,
      role: "user",
      timestamp: new Date(),
    };

    setMessages((prev) => [...prev, userMessage]);

    // Simulate assistant response (in a real app, this would be an API call)
    setTimeout(() => {
      const assistantMessage: Message = {
        id: uuidv4(),
        content: `I received your message: "${content}"`,
        role: "assistant",
        timestamp: new Date(),
      };

      setMessages((prev) => [...prev, assistantMessage]);
    }, 1000);
  };

  return (
    <div className="flex h-screen flex-col p-4">
      <Chat messages={messages} onSendMessage={handleSendMessage} />
    </div>
  );
}
