#!/usr/bin/env bash
set -e

APP_NAME="passive"
BUILD_DIR="build"

echo "🧹 Cleaning previous build..."
rm -rf "$BUILD_DIR"

echo "📁 Creating build directory..."
mkdir -p "$BUILD_DIR"

echo "☕ Building $APP_NAME..."
go build -o "$BUILD_DIR/$APP_NAME" ./cmd/passive

# echo "🚀 Running $APP_NAME..."
# echo "--------------------------------"
# "$BUILD_DIR/$APP_NAME" "$@"
# echo "--------------------------------"

echo "✅ Done."
