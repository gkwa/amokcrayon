#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <image_file>"
    exit 1
fi

IMAGE_FILE="$1"

# Check if the file exists
if [ ! -f "$IMAGE_FILE" ]; then
    echo "Error: File '$IMAGE_FILE' not found."
    exit 1
fi

# Ensure the file is writable
chmod u+w "$IMAGE_FILE"

# Add custom metadata using the config file
exiftool -config custom_config.config -overwrite_original \
  "-XMP-eden:Product=Sushi Nori" \
  "-XMP-eden:Brand=Eden" \
  "-XMP-eden:Weight=0.6 oz" \
  "-XMP-eden:Sheets=7" \
  "-XMP-eden:Price=8.69" \
  "-XMP-eden:IsGlutenFree=Yes" \
  "-XMP-eden:IsVegan=Yes" \
  "$IMAGE_FILE"

echo "Custom metadata added. Verifying:"
exiftool -config custom_config.config -G1 -s -XMP-eden:all "$IMAGE_FILE"

