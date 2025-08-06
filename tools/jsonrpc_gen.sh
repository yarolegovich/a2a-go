#!/bin/bash

set -euo pipefail

# Configuration
SCHEMA_URL="https://raw.githubusercontent.com/a2aproject/A2A/main/specification/json/a2a.json"
SCHEMA_FILE="a2a.json"
OUTPUT_DIR="./internal/jsonrpc"
OUTPUT_FILE="spec.go"
PACKAGE_NAME="jsonrpc"

mkdir -p "$OUTPUT_DIR"

echo "Created output directory: $OUTPUT_DIR"

echo "Downloading A2A JSON schema..."

curl -s -o "$SCHEMA_FILE" "$SCHEMA_URL"

echo "Generating Go types..."

go-jsonschema \
    --package "$PACKAGE_NAME" \
    --output "$OUTPUT_DIR/$OUTPUT_FILE" \
    "$SCHEMA_FILE"

echo "Formatting generated Go code..."

gofmt -w "$OUTPUT_DIR/$OUTPUT_FILE"

# Clean up downloaded schema file
rm -f "$SCHEMA_FILE"
echo "Cleaned up temporary files"

echo "Done: $OUTPUT_DIR/$OUTPUT_FILE"