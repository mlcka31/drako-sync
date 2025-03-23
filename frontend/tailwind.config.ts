import { type Config } from "tailwindcss";
import { fontFamily } from "tailwindcss/defaultTheme";

const config = {
  content: ["./src/**/*.tsx"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["var(--font-superstar)", ...fontFamily.sans],
        mono: ["var(--font-geist-mono)", ...fontFamily.mono],
        geist: ["var(--font-geist-sans)", ...fontFamily.sans],
        superstar: ["var(--font-superstar)", "sans-serif"],
        superstarOrig: ["var(--font-superstar-orig)", "sans-serif"],
      },
      colors: {
        // Primary colors
        primary: {
          DEFAULT: "#6366F1", // Indigo
          light: "#818CF8",
          dark: "#4F46E5",
        },
        // Background colors
        background: {
          DEFAULT: "#0F172A", // Dark blue
          light: "#1E293B",
          dark: "#020617",
        },
        // Text colors
        text: {
          DEFAULT: "#F8FAFC", // Light gray
          secondary: "#94A3B8",
          muted: "#64748B",
        },
        // Accent colors
        accent: {
          DEFAULT: "#10B981", // Emerald
          light: "#34D399",
          dark: "#059669",
        },
        // Alert/status colors
        error: "#EF4444",
        warning: "#F59E0B",
        success: "#10B981",
        info: "#3B82F6",
      },
      borderRadius: {
        xl: "1rem",
        "2xl": "1.5rem",
      },
      boxShadow: {
        glow: "0 0 20px rgba(99, 102, 241, 0.5)",
        card: "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
      },
      animation: {
        "pulse-slow": "pulse 3s cubic-bezier(0.4, 0, 0.6, 1) infinite",
      },
    },
  },
  plugins: [],
} satisfies Config;

module.exports = config;
