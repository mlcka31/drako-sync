"use client";

import "@rainbow-me/rainbowkit/styles.css";
import "~/styles/globals.css";

import { GeistSans } from "geist/font/sans";
import { type Metadata } from "next";
import Providers from "~/providers/Providers";
import { superstarFont, superstarOrigFont } from "./fonts";
import { useEffect } from "react";

// export const metadata: Metadata = {
//   title: "Drako Sync",
//   description: "Defeat the AI Agent to win the prize pool",
//   icons: [{ rel: "icon", url: "/favicon.ico" }],
// };

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  useEffect(() => {
    // Apply overscroll behavior to prevent pull-to-refresh
    document.documentElement.style.overscrollBehavior = "none";
    document.body.style.overscrollBehavior = "none";

    return () => {
      // Clean up
      document.documentElement.style.overscrollBehavior = "";
      document.body.style.overscrollBehavior = "";
    };
  }, []);

  return (
    <html
      lang="en"
      className={`${GeistSans.variable} ${superstarFont.variable} ${superstarOrigFont.variable}`}
    >
      <body className="h-[100dvh]">
        {/* Background images with responsive handling */}
        <div className="absolute inset-0 z-0 h-full w-full">
          <div className="hidden h-full w-full md:block">
            <img
              src="/assets/web-hero-image.png"
              alt="Background"
              className="h-full w-full object-cover"
            />
          </div>
          <div className="block h-full w-full md:hidden">
            <img
              src="/assets/mobile-hero-image.png"
              alt="Mobile Background"
              className="h-full w-full object-cover"
            />
          </div>
        </div>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
