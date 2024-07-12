#!/bin/bash

if [ $# -lt 3 ]; then
    echo "Usage: $0 <image_file> <tag> <value>"
    exit 1
fi

IMAGE_FILE="$1"
TAG="$2"
VALUE="$3"

if [ ! -f "$IMAGE_FILE" ]; then
    echo "Error: File not found."
    exit 1
fi

exiftool -overwrite_original "-$TAG=$VALUE" "$IMAGE_FILE"

echo "Metadata added. Verifying:"
exiftool -$TAG "$IMAGE_FILE"
