# Use the official Node.js image.
# https://hub.docker.com/_/node
FROM node:18-alpine

# Create and change to the app directory.
WORKDIR /usr/src/app

# Copy application dependency manifests to the container image.
# A wildcard is used to ensure both package.json AND package-lock.json are copied.
COPY package*.json ./

# Install production dependencies.
RUN npm install --only=production

# Copy local code to the container image.
COPY . .

# Build the TypeScript code.
RUN npm run build

# Expose the port that your application will listen on
EXPOSE 8080

# Run the web service on container startup.
CMD [ "node", "dist/index.js" ]
