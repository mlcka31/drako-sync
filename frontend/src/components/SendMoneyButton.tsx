"use client";

import { useSendTransaction } from "wagmi";
import { parseEther } from "viem";

const SendMoneyButton = () => {
  const { sendTransaction } = useSendTransaction();

  const handleClick = () => {
    // Replace with actual transaction details
    const transactionDetails = {
      to: "recipient_address_here",
      value: "amount_here",
    };

    sendTransaction({
      to: "0xd2135CfB216b74109775236E36d4b433F1DF507B",
      value: parseEther("0.01"),
    });
  };

  return <button onClick={handleClick}>Send Money</button>;
};

export default SendMoneyButton;
