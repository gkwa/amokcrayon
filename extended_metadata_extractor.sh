#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 <image_file>"
    exit 1
fi

IMAGE_FILE="$1"

if [ ! -f "$IMAGE_FILE" ]; then
    echo "Error: File not found."
    exit 1
fi

exiftool -a -u -g1 "$IMAGE_FILE"

identify -verbose "$IMAGE_FILE"

exiv2 -pa "$IMAGE_FILE"
