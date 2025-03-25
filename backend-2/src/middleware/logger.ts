import { FastifyRequest, FastifyReply } from "fastify";

// Create a logger object with basic methods
const logger = {
  info: (message: any) => console.log(JSON.stringify(message)),
  error: (message: any, error?: any) => console.error(message, error || ""),
  debug: (message: any) => console.log(`DEBUG: ${message}`),
};

// Middleware function for request logging
export async function loggerMiddleware(
  request: FastifyRequest,
  reply: FastifyReply
) {
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
export default logger;
