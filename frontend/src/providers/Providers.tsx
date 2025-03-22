"use client";

import { Chain, getDefaultConfig } from "@rainbow-me/rainbowkit";
import { cookieStorage, createStorage, http } from "wagmi";
import {
  blastSepolia,
  bscTestnet,
  sepolia,
  zksync,
  zksyncSepoliaTestnet,
} from "wagmi/chains";
import { RainbowKitProvider } from "@rainbow-me/rainbowkit";

import "~/styles/globals.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

import { WagmiProvider } from "wagmi";

type Props = {
  children?: React.ReactNode;
};

const projectId = "32e0c2400c7a431216898567498a882b";

const supportedChains: Chain[] = [sepolia, zksync, zksyncSepoliaTestnet];

export const config = getDefaultConfig({
  appName: "WalletConnection",
  projectId,
  chains: supportedChains as any,
  ssr: false,
  storage: createStorage({
    storage: cookieStorage,
  }),
  transports: supportedChains.reduce(
    (obj, chain) => ({ ...obj, [chain.id]: http() }),
    {},
  ),
});

const Providers = ({ children }: Props) => {
  return (
    <QueryClientProvider client={queryClient}>
      <WagmiProvider config={config}>
        <RainbowKitProvider modalSize="compact">{children}</RainbowKitProvider>
      </WagmiProvider>
    </QueryClientProvider>
  );
};

export default Providers;
