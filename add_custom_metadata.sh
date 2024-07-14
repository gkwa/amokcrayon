#!/bin/bash

CONFIG_FILE="custom_config.config.pl"

# Function to display help message
show_help() {
    cat <<EOF
Usage: $0 <image_file>
This script adds custom EXIF metadata to an image file.

Arguments:
 <image_file>    Path to the image file to process

Note: This script requires a config file named '$CONFIG_FILE' in the same
     directory.

Example:
 $0 path/to/your/image.jpg
EOF
}

# Check if help is requested
if [[ $1 == "--help" || $1 == "-h" ]]; then
    show_help
    exit 0
fi

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Error: No image file specified."
    echo
    show_help
    exit 1
fi

IMAGE_FILE="$1"

# Check if the image file exists
if [ ! -f "$IMAGE_FILE" ]; then
    echo "Error: File '$IMAGE_FILE' not found."
    exit 1
fi

# Check if the config file exists
if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: Config file '$CONFIG_FILE' not found."
    echo "Please ensure the config file is in the same directory as this script."
    exit 1
fi

# Ensure the file is writable
chmod u+w "$IMAGE_FILE"

# Add custom metadata using the config file
if exiftool -config "$CONFIG_FILE" -overwrite_original \
    "-XMP-eden:Product=Sushi Nori" \
    "-XMP-eden:Brand=Eden" \
    "-XMP-eden:Weight=0.6 oz" \
    "-XMP-eden:Sheets=7" \
    "-XMP-eden:Price=8.69" \
    "-XMP-eden:IsGlutenFree=Yes" \
    "-XMP-eden:IsVegan=Yes" \
    "$IMAGE_FILE"; then
    echo "Custom metadata added successfully. Verifying:"
    exiftool -config "$CONFIG_FILE" -G1 -s -XMP-eden:all "$IMAGE_FILE"
else
    echo "Error: Failed to add custom metadata. Please check your exiftool"
    echo "installation and config file."
fi
