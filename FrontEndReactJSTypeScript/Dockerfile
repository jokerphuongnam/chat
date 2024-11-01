# Use Node.js as the base image
FROM node:18 AS builder

WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY . .

# Build the React app
RUN npm run build

# Stage 2: Serve the app using a lightweight web server
FROM nginx:alpine

# Copy built assets from the previous stage
COPY --from=builder /app/build /usr/share/nginx/html

# Expose the port Nginx is running on
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]