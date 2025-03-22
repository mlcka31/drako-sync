"use client";

import React from "react";
import Link from "next/link";
import {
  backendApp,
  backendAttestation,
  contractAddress,
  githubRepo,
} from "~/config";
import { useGame } from "~/hooks/useGame";

export default function RulesPage() {
  const { agentAddress } = useGame();
  return (
    <div className="z-100 relative min-h-screen w-full bg-[#070716] p-4 text-[#f0f0f0]">
      <div className="mx-auto max-w-4xl py-8">
        <h1 className="mb-6 text-center text-4xl font-bold">
          Drako Sync: Game Rules
        </h1>

        <section className="mb-8">
          <h2 className="mb-4 text-2xl font-semibold">Project Overview</h2>
          <p className="mb-4">
            Drako Sync is a decentralized, AI-powered Dungeons & Dragons game
            where the Game Master is an autonomous AI running in a Trusted
            Execution Environment (TEE). This innovative approach combines
            blockchain technology, artificial intelligence, and classic
            role-playing gameplay to create a trustless and immersive gaming
            experience.
          </p>
        </section>

        <section className="mb-8">
          <h2 className="mb-4 text-2xl font-semibold">Key Features</h2>
          <ul className="list-disc space-y-2 pl-6">
            <li>
              <strong>AI Game Master:</strong> An advanced AI serves as the Game
              Master, creating adventures, narrating stories, and adjudicating
              gameplay decisions.
            </li>
            <li>
              <strong>Trusted Execution Environment:</strong> The AI runs in a
              TEE, ensuring that no one—not even the developers—can access its
              private keys or manipulate game outcomes.
            </li>
            <li>
              <strong>Blockchain Integration:</strong> Built on Ethereum, the
              game uses smart contracts to handle transactions and prize
              distribution.
            </li>
            <li>
              <strong>Pay-to-Play Mechanics:</strong> Players send messages to
              interact with the game world by paying a small fee in ETH.
            </li>
            <li>
              <strong>Escalating Stakes:</strong> Message costs increase with
              each game iteration, creating a growing prize pool.
            </li>
            <li>
              <strong>Winner-Takes-All:</strong> The player who successfully
              completes the adventure or achieves the game's objective wins the
              entire prize pool.
            </li>
          </ul>
        </section>

        <section className="mb-8">
          <h2 className="mb-4 text-2xl font-semibold">How It Works</h2>
          <ol className="list-decimal space-y-2 pl-6">
            <li>Players connect their Web3 wallet to join the game.</li>
            <li>Each player interaction requires a small ETH fee.</li>
            <li>
              The AI Game Master responds to player actions and guides the
              adventure.
            </li>
            <li>
              As the game progresses, message fees increase, growing the prize
              pool.
            </li>
            <li>
              The AI autonomously determines the winner based on gameplay.
            </li>
            <li>
              Smart contracts automatically distribute the prize to the winner's
              wallet.
            </li>
          </ol>
        </section>

        <section className="mb-8">
          <h2 className="mb-4 text-2xl font-semibold">Security & Fairness</h2>
          <p className="mb-4">
            The use of a Trusted Execution Environment ensures that the game is
            provably fair. The AI's decision-making process is secure and
            tamper-proof, guaranteeing that no one can influence the outcome of
            the game for their benefit.
          </p>
        </section>

        <section className="mb-8">
          <h2 className="mb-4 text-2xl font-semibold">Technology Stack</h2>
          <ul className="list-disc space-y-2 pl-6">
            <li>
              <strong>Frontend:</strong> React, Wagmi, RainbowKit
            </li>
            <li>
              <strong>Backend:</strong> GO, TEE infrastructure on Phala.network
            </li>
            <li>
              <strong>Blockchain:</strong> ZSync, Solidity smart contracts
            </li>
            <li>
              <strong>AI:</strong> Advanced language model running in a secure
              enclave
            </li>
          </ul>
        </section>

        <div className="mt-10 flex flex-col items-center gap-4">
          <Link href="/" className="text-blue-400 hover:text-blue-300">
            Start Game
          </Link>
          <Link
            href={githubRepo}
            className="flex flex-row items-center gap-2 text-blue-400 hover:text-blue-300"
          >
            <img src="/assets/github.svg" alt="Github" width={24} height={24} />
            Github
          </Link>
          <Link
            href={backendAttestation}
            className="text-blue-400 hover:text-blue-300"
          >
            Backend attestation
          </Link>
          <Link href={backendApp} className="text-blue-400 hover:text-blue-300">
            Agent TEE
          </Link>
          <Link href={"#"} className="text-blue-400 hover:text-blue-300">
            Contract Address {contractAddress}
          </Link>
          <Link href={"#"} className="text-blue-400 hover:text-blue-300">
            Agent Address: {agentAddress}
          </Link>
        </div>
      </div>
    </div>
  );
}
