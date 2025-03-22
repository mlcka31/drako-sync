"use client";

import { Chain, getDefaultConfig } from "@rainbow-me/rainbowkit";
import { cookieStorage, createStorage, http } from "wagmi";
import { blastSepolia, bscTestnet, sepolia } from "wagmi/chains";

// import { SessionProvider } from "next-auth/react";
import "~/styles/globals.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

// return <QueryClientProvider client={queryClient}>...</QueryClientProvider>;

import { WagmiProvider } from "wagmi";

type Props = {
  children?: React.ReactNode;
};

const projectId = "32e0c2400c7a431216898567498a882b";

const supportedChains: Chain[] = [sepolia, bscTestnet, blastSepolia];

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

// export const NextAuthProvider = ({ children }: Props) => {
//   return <SessionProvider>{children}</SessionProvider>;
// };

const Providers = ({ children }: Props) => {
  return (
    // <NextAuthProvider>
    <QueryClientProvider client={queryClient}>
      <WagmiProvider config={config}>{children}</WagmiProvider>
    </QueryClientProvider>
    // </NextAuthProvider>
  );
};

export default Providers;
