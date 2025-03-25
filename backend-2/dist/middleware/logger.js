"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.loggerMiddleware = loggerMiddleware;
// Create a logger object with basic methods
const logger = {
    info: (message) => console.log(JSON.stringify(message)),
    error: (message, error) => console.error(message, error || ""),
    debug: (message) => console.log(`DEBUG: ${message}`),
};
// Middleware function for request logging
async function loggerMiddleware(request, reply) {
    const start = Date.now();
    // Use the onResponse hook which is the proper way in Fastify
    reply.raw.on("finish", () => {
        const duration = Date.now() - start;
        logger.info({
            url: request.url,
            method: request.method,
            statusCode: reply.statusCode,
            duration: `${duration}ms`,
        });
    });
}
// Export the logger for use in other files
exports.default = logger;
