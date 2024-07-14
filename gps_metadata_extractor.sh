#!/bin/bash

# Function to display help message
show_help() {
    echo "Usage: $0 <image_file>"
    echo
    echo "This script extracts GPS coordinates from the EXIF data of an image file"
    echo "and provides a Google Maps link to the location."
    echo
    echo "Arguments:"
    echo "  <image_file>    Path to the image file to process"
    echo
    echo "Example:"
    echo "  $0 path/to/your/image.jpg"
}

# Function to extract coordinates
extract_coords() {
   local lat=$(exiftool -n -p '$GPSLatitude' "$1" 2>/dev/null)
   local lon=$(exiftool -n -p '$GPSLongitude' "$1" 2>/dev/null)
   local lat_ref=$(exiftool -n -p '$GPSLatitudeRef' "$1" 2>/dev/null)
   local lon_ref=$(exiftool -n -p '$GPSLongitudeRef' "$1" 2>/dev/null)
   if [ -n "$lat" ] && [ -n "$lon" ]; then
       [ "$lat_ref" = "S" ] && lat="-$lat"
       [ "$lon_ref" = "W" ] && lon="-$lon"
       lon="${lon#-}"
       echo "GPS Coordinates found!"
       echo "Google Maps Link: https://www.google.com/maps?q=${lat},${lon}"
   else
       echo "No GPS coordinates found in the image."
   fi
}

# Check if help is requested
if [[ "$1" == "--help" || "$1" == "-h" ]]; then
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

# Check if the file exists
if [ ! -f "$IMAGE_FILE" ]; then
    echo "Error: File '$IMAGE_FILE' not found."
    exit 1
fi

# Process the image file
extract_coords "$IMAGE_FILE"