package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	packages := []string{
		"exiv2",
		"gdal-bin",
		"imagemagick",
		"jhead",
		"libimage-exiftool-perl",
	}

	ubuntu := client.Container().From("ubuntu:22.04")

	// Create a cache volume
	cache := client.CacheVolume("apt-cache")

	// Install packages with caching
	container := ubuntu.
		WithMountedCache("/var/cache/apt", cache).
		WithExec([]string{"apt-get", "update"}).
		WithExec(append([]string{"apt-get", "install", "--assume-yes"}, packages...))

	// Create a separate step for version checking
	versionContainer := container.WithExec([]string{"/bin/sh", "-c", `
		exiftool -ver &&
		identify -version | head -n 1 &&
		jhead -V &&
		exiv2 -V &&
		gdalinfo --version
	`})

	output, err := versionContainer.Stdout(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
