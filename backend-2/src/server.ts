import Fastify from "fastify";
import cors from "@fastify/cors";
import dotenv from "dotenv";
import apiRoutes from "./routes/api";
import logger from "./middleware/logger";
import { blockchainService } from "./services/blockchainService";
import { lookup } from "dns";
import { walletService } from "./services/walletService";

// Load environment variables
dotenv.config();

const server = Fastify({
  logger: false,
});

// Register CORS
server.register(cors, {
  origin: true,
});

// Register API routes
server.register(apiRoutes, { prefix: "/api" });

// Add public key route
server.get("/public-key", async (request, reply) => {
  try {
    const publicKey = walletService.getWallet().publicKey;
    return { publicKey };
  } catch (error) {
    logger.error("Error retrieving public key:", error);
    return reply.code(500).send({ error: "Failed to retrieve public key" });
  }
});

// Start the server
const start = async () => {
  try {
    const PORT = process.env.PORT || 3000;
    const HOST = process.env.HOST || "0.0.0.0";

    await server.listen({ port: Number(PORT), host: HOST });
    logger.info(`Server listening on ${HOST}:${PORT}`);

    // Auto-initialize blockchain service if contract address is provided in env
    if (!process.env.CONTRACT_ADDRESS)
      throw new Error("CONTRACT_ADDRESS is not set");
    const initialized = await blockchainService.init(
      process.env.CONTRACT_ADDRESS
    );
    if (initialized) {
      blockchainService.startPolling(
        Number(process.env.POLLING_INTERVAL) || 10000
      );
      logger.info(
        `Blockchain polling started automatically for contract: ${process.env.CONTRACT_ADDRESS}`
      );
    }
  } catch (err) {
    logger.error("Error starting server:", err);
    process.exit(1);
  }
};

start();

export default server;
