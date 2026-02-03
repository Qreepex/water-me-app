#!/bin/bash

# Build script for Water Me backend components

set -e

# Read version from VERSION file, allow override via argument
VERSION=${1:-$(cat VERSION)}
REGISTRY="ghcr.io/qreepex"

echo "Building Water Me backend components v${VERSION}"

# Build API server
echo "Building API server..."
docker build --build-arg COMPONENT=api \
  -t ${REGISTRY}/plants-backend-api:${VERSION} \
  -t ${REGISTRY}/plants-backend-api:latest \
  .

# Build notification worker
echo "Building notification worker..."
docker build --build-arg COMPONENT=notification-worker \
  -t ${REGISTRY}/plants-notification-worker:${VERSION} \
  -t ${REGISTRY}/plants-notification-worker:latest \
  .

echo "Build complete!"
echo ""
echo "To push to registry, run:"
echo "  docker push ${REGISTRY}/plants-backend-api:${VERSION}"
echo "  docker push ${REGISTRY}/plants-backend-api:latest"
echo "  docker push ${REGISTRY}/plants-notification-worker:${VERSION}"
echo "  docker push ${REGISTRY}/plants-notification-worker:latest"
