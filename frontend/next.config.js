/**
 * Run `build` or `dev` with `SKIP_ENV_VALIDATION` to skip env validation. This is especially useful
 * for Docker builds.
 */
import "./src/env.js";

/** @type {import("next").NextConfig} */
const config = {
  // Make sure experimental features are enabled if needed
  experimental: {
    // ... any experimental features ...
  },

  // Add any required redirects
  async redirects() {
    return [
      // You can add redirects here if needed
    ];
  },
};

export default config;
