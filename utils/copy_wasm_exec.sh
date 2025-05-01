
#!/bin/bash

set -e

# Find the Go root directory
GOROOT=$(go env GOROOT)

# Find wasm_exec.js
WASM_EXEC="$GOROOT/misc/wasm/wasm_exec.js"

# Check if it exists
if [ ! -f "$WASM_EXEC" ]; then
    echo "Error: wasm_exec.js not found in $WASM_EXEC"
    exit 1
fi

# Copy it to project root
cp "$WASM_EXEC" ./wasm_exec.js

echo "✅ wasm_exec.js copied from $WASM_EXEC to project root"

