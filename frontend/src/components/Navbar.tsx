import React from "react";
import Logo from "./Logo";
import { ConnectBtn } from "./ConnectButton";
import Link from "next/link";

const Navbar = () => {
  return (
    <nav className="fixed left-0 right-0 top-0 z-50 flex items-center justify-between px-4 py-3 md:px-8 md:py-4">
      <div className="flex items-center">
        <Logo className="h-8 w-auto md:h-10" />
      </div>
      <div className="flex items-center gap-4">
        <Link
          href="/about"
          className="hidden text-white hover:text-gray-300 md:block"
        >
          About
        </Link>
        <Link
          href="/rules"
          className="hidden text-white hover:text-gray-300 md:block"
        >
          Rules
        </Link>
        <ConnectBtn />
      </div>
    </nav>
  );
};

export default Navbar;
