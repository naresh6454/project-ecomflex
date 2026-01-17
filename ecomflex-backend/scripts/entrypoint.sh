#!/bin/sh
# File: ecomflex-backend/scripts/entrypoint.sh

set -e

echo "Starting application initialization..."

# Run migrations
echo "Running database migrations..."
/app/ecomflex-api migrate

# Start the application
echo "Starting API server..."
exec /app/ecomflex-api