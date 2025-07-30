#!/bin/bash
set -x

# Source directory (default: current directory)
SRC_DIR="${1:-.}"

# Output directory
OUT_DIR="./utf8_output"
mkdir -p "$OUT_DIR"

# Ensure source directory is absolute
SRC_DIR_ABS=$(cd "$SRC_DIR"; pwd)
OUT_DIR_ABS=$(cd "$OUT_DIR"; pwd)

# Find all .txt files in the source directory and subdirectories
find "$SRC_DIR_ABS" -type f | while read -r FILE; do
    # Get relative path from source directory
    REL_PATH="${FILE#$SRC_DIR_ABS/}"

    # Construct output file path
    OUT_FILE="$OUT_DIR_ABS/$REL_PATH"

    # Make sure the output directory exists
    mkdir -p "$(dirname "$OUT_FILE")"

    echo "Converting: $FILE → $OUT_FILE"

    if iconv -f GB2312 -t UTF-8 "$FILE" > "$OUT_FILE"; then
        echo "✅ Success"
    else
        echo "⚠️ Failed to convert: $FILE"
    fi
done

echo "✅ All files converted. Output in: $OUT_DIR"

