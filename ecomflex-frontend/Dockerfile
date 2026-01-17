# Multi-stage build for ecomflex-frontend
# Stage 1: Build with Node
FROM node:18-alpine AS builder

# Accept build-time API URL, but let .env file take precedence
ARG VITE_API_BASE_URL

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Only set environment variable if ARG was provided, otherwise use .env file
RUN if [ -n "$VITE_API_BASE_URL" ]; then export VITE_API_BASE_URL=$VITE_API_BASE_URL; fi

# Build the application
RUN npm run build

# Stage 2: Serve with nginx
FROM nginx:alpine

# Remove default nginx config
RUN rm /etc/nginx/conf.d/default.conf

# Copy our nginx config
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built app from builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]