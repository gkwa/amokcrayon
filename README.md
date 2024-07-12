# GPS Tools Installer Dagger Pipeline

## Prerequisites
1. Install Go (1.20 or later)
2. Install Docker

## Setup and Run
1. Create a new directory for your project
2. Save the `main.go` file in this directory
3. Initialize the Go module:
   go mod init gps_tools_installer
4. Add the Dagger dependency:
   go get dagger.io/dagger@v0.9.3
5. Run the Dagger pipeline:
   go run main.go

This will execute the pipeline, which will create a container, install the specified tools, and print the version information for each tool.

## Troubleshooting
- Ensure Docker is running on your system
- If you encounter any permission issues, you may need to run the command with sudo:
  sudo go run main.go

Note: The first run may take some time as it downloads the necessary container images and installs the tools.
