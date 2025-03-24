"use client";

import Link from "next/link";
import { ConnectBtn } from "./ConnectButton";
import Logo from "./Logo";

const Navbar = () => {
  return (
    <nav
      className={`fixed left-0 right-0 top-0 z-50 flex w-full flex-col items-center justify-between bg-[#070716] px-4 py-3 md:flex-row md:px-8 md:py-4`}
    >
      <Logo className="h-8 w-auto md:h-10" />
      <div className="flex items-center gap-4">
        <Link href="/rules" className="text-white hover:text-gray-300 md:block">
          Rules
        </Link>
        <ConnectBtn />
      </div>
    </nav>
  );
};

export default Navbar;
