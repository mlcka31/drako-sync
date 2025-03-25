"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fastify_1 = __importDefault(require("fastify"));
const cors_1 = __importDefault(require("@fastify/cors"));
const dotenv_1 = __importDefault(require("dotenv"));
const api_1 = __importDefault(require("./routes/api"));
const logger_1 = __importDefault(require("./middleware/logger"));
const blockchainService_1 = require("./services/blockchainService");
const walletService_1 = require("./services/walletService");
// Load environment variables
dotenv_1.default.config();
const server = (0, fastify_1.default)({
    logger: false,
});
// Register CORS
server.register(cors_1.default, {
    origin: true,
});
// Register API routes
server.register(api_1.default, { prefix: "/api" });
// Add public key route
server.get("/public-key", async (request, reply) => {
    try {
        const publicKey = walletService_1.walletService.getWallet().publicKey;
        return { publicKey };
    }
    catch (error) {
        logger_1.default.error("Error retrieving public key:", error);
        return reply.code(500).send({ error: "Failed to retrieve public key" });
    }
});
// Start the server
const start = async () => {
    try {
        const PORT = process.env.PORT || 3000;
        const HOST = process.env.HOST || "0.0.0.0";
        await server.listen({ port: Number(PORT), host: HOST });
        logger_1.default.info(`Server listening on ${HOST}:${PORT}`);
        // Auto-initialize blockchain service if contract address is provided in env
        if (!process.env.CONTRACT_ADDRESS)
            throw new Error("CONTRACT_ADDRESS is not set");
        const initialized = await blockchainService_1.blockchainService.init(process.env.CONTRACT_ADDRESS);
        if (initialized) {
            blockchainService_1.blockchainService.startPolling(Number(process.env.POLLING_INTERVAL) || 10000);
            logger_1.default.info(`Blockchain polling started automatically for contract: ${process.env.CONTRACT_ADDRESS}`);
        }
    }
    catch (err) {
        logger_1.default.error("Error starting server:", err);
        process.exit(1);
    }
};
start();
exports.default = server;
