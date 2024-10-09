#!/bin/bash

# Define the output binary name prefix
BINARY_NAME="dankey"

# Define the platforms and architectures
platforms=("linux" "windows")
archs=("amd64" "arm64")

# Check if UPX is installed
if ! command -v upx &> /dev/null; then
    echo "UPX is not installed. Please install UPX to compress binaries."
    exit 1
fi

# Loop through each platform and architecture
for platform in "${platforms[@]}"; do
  for arch in "${archs[@]}"; do
    # Set the output file name
    output_name="${BINARY_NAME}_${platform}_${arch}"
    
    # Add .exe extension for Windows binaries
    if [ "$platform" == "windows" ]; then
      output_name="${output_name}.exe"
    fi

    # Build the binary
    echo "Building for $platform $arch..."
    GOOS=$platform GOARCH=$arch go build -ldflags "-s -w" -o "$output_name"

    # Check if the build was successful
    if [ $? -ne 0 ]; then
      echo "Failed to build for $platform $arch"
      continue
    fi

    # Compress the binary with UPX
    echo "Compressing $output_name with UPX..."
    upx "$output_name"

    # Check if UPX compression was successful
    if [ $? -ne 0 ]; then
      echo "Failed to compress $output_name with UPX"
    else
      echo "Successfully built and compressed $output_name"
    fi
  done
done

