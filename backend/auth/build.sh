#!/bin/bash

set -e

echo "=== Building auth service ==="

# 编译Go代码
echo "Compiling..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth.exe .

echo "Build completed: auth.exe"

# 构建Docker镜像
echo "Building Docker image..."
docker build -t auth:latest .

echo "=== Docker image built successfully ==="
