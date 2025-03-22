import React from "react";
import { cva, type VariantProps } from "class-variance-authority";

const buttonVariants = cva(
  "flex items-center justify-center rounded-lg font-semibold transition-all duration-200 ease-in-out",
  {
    variants: {
      variant: {
        primary:
          "bg-gradient-to-r from-blue-500 to-blue-600 text-white hover:from-blue-600 hover:to-blue-700 hover:shadow-md hover:shadow-blue-500/30 active:shadow-sm active:shadow-blue-500/20",
        secondary:
          "bg-slate-800 text-white hover:bg-slate-900 hover:shadow-md hover:shadow-blue-500/30 active:shadow-sm active:shadow-blue-500/20",
        outline:
          "bg-transparent text-blue-500 border-2 border-blue-500 hover:bg-blue-500/10 hover:shadow-md hover:shadow-blue-500/30 active:shadow-sm active:shadow-blue-500/20",
      },
      size: {
        small: "text-sm py-2 px-4",
        medium: "text-base py-3 px-5",
        large: "text-lg py-4 px-6",
      },
      fullWidth: {
        true: "w-full",
        false: "w-auto",
      },
    },
    defaultVariants: {
      variant: "primary",
      size: "medium",
      fullWidth: false,
    },
  },
);

export interface FigmaButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {
  children: React.ReactNode;
}

const FigmaButton: React.FC<FigmaButtonProps> = ({
  className,
  variant,
  size,
  fullWidth,
  children,
  ...props
}) => {
  return (
    <button className={`rounded-lg bg-[#00F2FF33] px-3 py-2`} {...props}>
      {children}
    </button>
  );
};

export default FigmaButton;
