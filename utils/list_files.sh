#!/bin/bash
# Usage: ./list_source.sh [directory]
# Defaults to the current directory if none is specified.

DIR="${1:-.}"

# Ensure the directory exists
if [[ ! -d "$DIR" ]]; then
    echo "Error: Directory '$DIR' does not exist."
    exit 1
fi

# Find .go and .html files while skipping hidden directories
find "$DIR" -type d -name ".*" -o -type f \( -iname "*.go" -o -iname "*.html" -o -iname "*.md" \) -print0 | while IFS= read -r -d '' file; do
    echo -e "\n\nFilename: $file"
    echo "-----------------------"
    cat "$file"
    echo -e "\n\n\n"  # Three blank lines as a separator
done
