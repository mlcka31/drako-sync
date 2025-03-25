"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.apiRoutes = void 0;
exports.default = apiRoutes;
// Define API routes
async function apiRoutes(fastify) {
    // Health check endpoint
    fastify.get("/health", async () => {
        return { status: "ok" };
    });
}
// Export as named export for compatibility with current import
var api_1 = require("./api");
Object.defineProperty(exports, "apiRoutes", { enumerable: true, get: function () { return __importDefault(api_1).default; } });
