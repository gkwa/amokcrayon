#!/bin/bash

IMAGE_FILE="20240708_193150.jpg"

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
