import { FastifyInstance } from "fastify";

// Define API routes
export default async function apiRoutes(fastify: FastifyInstance) {
  // Health check endpoint
  fastify.get("/health", async () => {
    return { status: "ok" };
  });
}

// Export as named export for compatibility with current import
export { default as apiRoutes } from "./api";
