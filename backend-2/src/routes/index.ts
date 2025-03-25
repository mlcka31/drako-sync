import { FastifyInstance } from "fastify";

// Register all routes
export default async function routes(fastify: FastifyInstance) {
  // Register other route modules
  await fastify.register(require("./api").default, { prefix: "/api" });
}
