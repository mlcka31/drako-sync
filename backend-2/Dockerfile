# Use Node.js LTS version as the base image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install dependencies
RUN yarn

# Copy the rest of the application code
COPY . .

# Build the application (if TypeScript is used)
RUN yarn build

# Expose the port your Fastify app runs on
# (adjust this if your app uses a different port)
EXPOSE 8080 

# Command to run the application
CMD ["yarn", "start"]
