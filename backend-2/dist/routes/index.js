"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = routes;
// Register all routes
async function routes(fastify) {
    // Register other route modules
    await fastify.register(require("./api").default, { prefix: "/api" });
}
