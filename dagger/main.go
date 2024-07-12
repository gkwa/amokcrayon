package main

import (
   "context"
   "dagger/amokcrayon/internal/dagger"
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

func (m *Amokcrayon) SetupImageContainer(ctx context.Context) *dagger.Container {
   return dag.Container().
   	From("alpine:latest").
   	WithExec([]string{"apk", "add", "--no-cache", "curl"}).
   	WithExec([]string{"curl", "-LO", "https://i.pinimg.com/564x/0a/b4/29/0ab42953f223f2934989955b58ec1b4b.jpg"})
}

func (m *Amokcrayon) InstallGPSTools(ctx context.Context) (string, error) {
   setupContainer := m.SetupImageContainer(ctx)
   
   container := dag.Container().
   	From("ubuntu:22.04").
   	WithExec([]string{"apt-get", "update"}).
   	WithExec([]string{"apt-get", "install", "-y", "libimage-exiftool-perl", "imagemagick", "jhead", "exiv2", "gdal-bin"}).
   	WithFile("/image.jpg", setupContainer.File("/0ab42953f223f2934989955b58ec1b4b.jpg"))

   return container.WithExec([]string{"/bin/sh", "-c", `
   exiftool -ver &&
   identify -version | head -n 1 &&
   jhead -V &&
   exiv2 -V &&
   gdalinfo --version &&
   cat /etc/os-release
   `}).Stdout(ctx)
}

func (m *Amokcrayon) GetImageSize(ctx context.Context) (string, error) {
   setupContainer := m.SetupImageContainer(ctx)
   
   return setupContainer.
   	WithExec([]string{"sh", "-c", "ls -l /0ab42953f223f2934989955b58ec1b4b.jpg | awk '{print $5}'"}).
   	Stdout(ctx)
}
