// A generated module for Amokcrayon functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"

	"dagger/amokcrayon/internal/dagger"
)

type Amokcrayon struct{}

func (m *Amokcrayon) Echo(stringArg string) string {
	return stringArg
}

// Returns a container that echoes whatever string argument is provided
func (m *Amokcrayon) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Amokcrayon) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

// InstallGPSTools installs GPS tools and returns version information
func (m *Amokcrayon) InstallGPSTools(ctx context.Context) (string, error) {
	container := dag.Container().
		From("ubuntu:22.04").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "libimage-exiftool-perl", "imagemagick", "jhead", "exiv2", "gdal-bin"})

	return container.WithExec([]string{"/bin/sh", "-c", `
   	exiftool -ver &&
   	identify -version | head -n 1 &&
   	jhead -V &&
   	exiv2 -V &&
   	gdalinfo --version &&
   	cat /etc/os-release
   `}).Stdout(ctx)
}
