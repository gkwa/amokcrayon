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

	ubuntu := client.Container().From("ubuntu:latest")

	container := ubuntu.WithExec([]string{"apt", "update"})

	container = container.WithExec(append([]string{"apt", "install", "--assume-yes"}, packages...))

	output, err := container.WithExec([]string{"/bin/sh", "-c", `
		exiftool -ver &&
		identify -version | head -n 1 &&
		jhead -V &&
		exiv2 -V &&
		gdalinfo --version
	`}).Stdout(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
