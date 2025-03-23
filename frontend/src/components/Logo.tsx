import Image from "next/image";

interface LogoProps {
  className?: string;
  width?: number;
  height?: number;
}

export default function Logo({
  className,
  width = 120,
  height = 40,
}: LogoProps) {
  return (
    <Image
      src="/assets/logo.svg"
      alt="Logo"
      width={width}
      height={height}
      className={className}
      priority
    />
  );
}
