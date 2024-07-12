package main

import (
	"context"
	"dagger/amokcrayon/internal/dagger"
	"fmt"
	"strings"
)

type Amokcrayon struct{}

func (m *Amokcrayon) Echo(stringArg string) string {
	return stringArg
}

func (m *Amokcrayon) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

func (m *Amokcrayon) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (m *Amokcrayon) SetupImageContainer(ctx context.Context, imageUrl string) *dagger.Container {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "add", "--no-cache", "curl"}).
		WithExec([]string{"curl", "-L", "-o", "image.jpg", imageUrl})
}

func (m *Amokcrayon) InstallGPSTools(ctx context.Context, imageUrl string) (*dagger.Container, error) {
	setupContainer := m.SetupImageContainer(ctx, imageUrl)
	scriptContent := `#!/bin/bash
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
extract_coords "$IMAGE_FILE"`

	packages := []string{
		"libimage-exiftool-perl",
		"imagemagick",
		"jhead",
		"exiv2",
		"gdal-bin",
	}

	container := dag.Container().
		From("ubuntu:22.04").
		WithExec([]string{"apt-get", "update"}).
		WithExec(append([]string{"apt-get", "install", "-y"}, packages...)).
		WithFile("/image.jpg", setupContainer.File("/image.jpg")).
		WithNewFile("/src/gps_metadata_extractor.sh", scriptContent, dagger.ContainerWithNewFileOpts{
			Permissions: 0755,
		})

	_, err := container.WithExec([]string{"/bin/sh", "-c", `
   exiftool -ver &&
   identify -version | head -n 1 &&
   jhead -V &&
   exiv2 -V &&
   gdalinfo --version &&
   cat /etc/os-release
   `}).Stdout(ctx)

	if err != nil {
		return nil, err
	}

	return container, nil
}

func (m *Amokcrayon) ExtractGPSMetadata(ctx context.Context) (string, error) {
	imageUrls := []string{
		"https://images.ctfassets.net/23aumh6u8s0i/m8orLPzuWhbaE35wCMNtr/194a52203ea7ccfc210f77fc0430410f/palm-tree-1",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/7/70/Nationalmuseum-karta.png/600px-Nationalmuseum-karta.png",
	}

	var results []string

	for _, url := range imageUrls {
		container, err := m.InstallGPSTools(ctx, url)
		if err != nil {
			return "", fmt.Errorf("failed to install GPS tools: %w", err)
		}

		output, err := container.WithExec([]string{"sh", "/src/gps_metadata_extractor.sh", "/image.jpg"}).Stdout(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to extract GPS metadata: %w", err)
		}

		results = append(results, fmt.Sprintf("Image URL: %s\nResult: %s\n", url, output))
	}

	return strings.Join(results, "\n"), nil
}

func (m *Amokcrayon) ExtractLocalGPSMetadata(ctx context.Context) (string, error) {
	scriptContent := `#!/bin/bash
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
extract_coords "$IMAGE_FILE"`

	packages := []string{
		"libimage-exiftool-perl",
		"imagemagick",
		"jhead",
		"exiv2",
		"gdal-bin",
	}

	container := dag.Container().
		From("ubuntu:22.04").
		WithExec([]string{"apt-get", "update"}).
		WithExec(append([]string{"apt-get", "install", "-y"}, packages...)).
		WithMountedDirectory("/src", dag.Directory()).
		WithNewFile("/src/gps_metadata_extractor.sh", scriptContent, dagger.ContainerWithNewFileOpts{
			Permissions: 0755,
		})

	output, err := container.WithExec([]string{"sh", "/src/gps_metadata_extractor.sh", "/src/20240712_103259.jpg"}).Stdout(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to extract GPS metadata: %w", err)
	}

	return fmt.Sprintf("Local Image: 20240712_103259.jpg\nResult: %s\n", output), nil
}
