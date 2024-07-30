#!/bin/env bash

APP_NAME="app"
OUTPUT_DIR="build"

mkdir -p $OUTPUT_DIR

build_app() {
  local os=$1
  local arch=$2
  local output_name="$OUTPUT_DIR/${APP_NAME}_${os}_${arch}"
  
  echo "Building for OS: $os, ARCH: $arch"
  GOOS=$os GOARCH=$arch go build -o $output_name .
  
  if [ $? -eq 0 ]; then
    echo "Build successful: ${output_name}_${os}_${arch}"
  else
    echo "Build failed for OS: $os, ARCH: $arch"
  fi
}

build_app "darwin" "amd64"

build_app "linux" "amd64"

build_app "windows" "amd64"

echo "Builds completed!"
