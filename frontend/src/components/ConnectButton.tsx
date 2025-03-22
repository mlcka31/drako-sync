"use client";

import { useEffect, useRef } from "react";
import {
  useConnectModal,
  useAccountModal,
  useChainModal,
} from "@rainbow-me/rainbowkit";
import { useAccount, useDisconnect } from "wagmi";
import FigmaButton from './FigmaButton';

// Helper function to shorten Ethereum addresses
const shortenAddress = (address?: string): string => {
  if (!address) return '';
  return `${address.slice(0, 6)}...${address.slice(-4)}`;
};

export const ConnectBtn = () => {
  const { isConnecting, address, isConnected, chain } = useAccount();
  const { openConnectModal } = useConnectModal();
  const { openAccountModal } = useAccountModal();
  const { openChainModal } = useChainModal();
  const { disconnect } = useDisconnect();

  const isMounted = useRef(false);

  useEffect(() => {
    isMounted.current = true;
  }, []);

  if (isConnected) {
    return (
      <FigmaButton 
        variant="outline" 
        onClick={() => disconnect()}
      >
        Disconnect {shortenAddress(address)}
      </FigmaButton>
    );
  }

  return (
    <FigmaButton 
      variant="primary" 
      onClick={() => openConnectModal && openConnectModal()}
      disabled={isConnecting}
    >
      {isConnecting ? "Connecting..." : "Connect Wallet"}
    </FigmaButton>
  );
};
