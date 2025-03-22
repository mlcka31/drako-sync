import React from "react";

import Link from "next/link";
import Logo from "~/components/Logo";

export default function StartPage() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-black p-4 text-white">
      <div className="flex w-full max-w-md flex-col items-center">
        {/* Logo */}
        <div className="mb-8">
          <Logo />
        </div>

        {/* Heading */}
        <h1 className="mb-6 text-center text-3xl font-bold">
          Welcome to the Hack
        </h1>

        {/* Description */}
        <p className="mb-10 text-center text-gray-300">
          Join us for an exciting hackathon experience where innovation meets
          collaboration.
        </p>

        {/* Buttons */}
        <div className="w-full space-y-4">
          <Link href="/register" className="block w-full">
            <button className="w-full rounded-lg bg-blue-600 px-4 py-3 font-medium text-white transition duration-200 hover:bg-blue-700">
              Register Now
            </button>
          </Link>

          <Link href="/login" className="block w-full">
            <button className="w-full rounded-lg border border-gray-600 bg-transparent px-4 py-3 font-medium text-white transition duration-200 hover:bg-gray-800">
              Log In
            </button>
          </Link>
        </div>

        {/* Footer text */}
        <p className="mt-12 text-center text-sm text-gray-400">
          By continuing, you agree to our Terms of Service and Privacy Policy.
        </p>
      </div>
    </div>
  );
}
