#!/bin/bash

sudo apt update
sudo apt install -y libimage-exiftool-perl imagemagick jhead exiv2 gdal-bin

exiftool -ver
identify -version | head -n 1
jhead -V
exiv2 -V
gdalinfo --version
