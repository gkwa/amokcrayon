#!/bin/bash
extract_coords() {
   local lat=$(exiftool -n -p '$GPSLatitude' "$1" 2>/dev/null)
   local lon=$(exiftool -n -p '$GPSLongitude' "$1" 2>/dev/null)
   local lat_ref=$(exiftool -n -p '$GPSLatitudeRef' "$1" 2>/dev/null)
   local lon_ref=$(exiftool -n -p '$GPSLongitudeRef' "$1" 2>/dev/null)
   if [ -n "$lat" ] && [ -n "$lon" ]; then
       [ "$lat_ref" = "S" ] && lat="-$lat"
       [ "$lon_ref" = "W" ] && lon="-$lon"
       lon="${lon#-}"
       echo "https://www.google.com/maps?q=${lat},${lon}"
   else
       echo "No GPS coordinates found."
   fi
}
if [ $# -eq 0 ]; then
   echo "Usage: $0 <image_file>"
   exit 1
fi
IMAGE_FILE="$1"
if [ ! -f "$IMAGE_FILE" ]; then
   echo "Error: File not found."
   exit 1
fi
extract_coords "$IMAGE_FILE"
