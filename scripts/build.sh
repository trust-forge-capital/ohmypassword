#!/bin/bash

set -e

echo "Building ohmypassword..."

# Get the directory of the script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

cd "$PROJECT_ROOT"

# Create output directory
mkdir -p bin

# Get version info
VERSION=${VERSION:-"dev"}
BUILD_TIME=${BUILD_TIME:-$(date -u '+%Y-%m-%d %H:%M:%S')}
GIT_COMMIT=${GIT_COMMIT:-$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")}

# Build for current platform
echo "Building for current platform..."
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
OUTPUT="bin/ohmypassword-${GOOS}-${GOARCH}"

if [ "$GOOS" = "windows" ]; then
    OUTPUT="${OUTPUT}.exe"
fi

LDFLAGS="-X main.version=$VERSION -X main.buildTime=$BUILD_TIME -X main.gitCommit=$GIT_COMMIT"

go build -ldflags="$LDFLAGS" -o "$OUTPUT" ./cmd/cli

echo "Built: $OUTPUT"

# Make executable
chmod +x "$OUTPUT"

echo "Done!"