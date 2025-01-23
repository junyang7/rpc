#!/usr/bin/env bash

set -e

ROOT=$(cd "$(dirname "$0")";pwd)
PROTO_DIR=${ROOT}"/put_your_proto_files_here"
DIST_GO=${ROOT}/dist/go
DIST_PHP=${ROOT}/dist/php

export https_proxy=http://127.0.0.1:7897 http_proxy=http://127.0.0.1:7897 all_proxy=socks5://127.0.0.1:7897

echo "开始检查 protoc 版本..."
# protoc 版本检查
protoc --version
if [ $? -ne 0 ]; then
  echo "protoc 未安装，正在安装..."
  brew install protobuf
else
  echo "protoc 已安装"
fi

echo "检查 grpc_php_plugin..."
# grpc_php_plugin 检查
if [ ! -f /usr/local/bin/grpc_php_plugin ]; then
  echo "grpc_php_plugin 未安装，正在安装..."
  brew install cmake protobuf
  git clone https://github.com/grpc/grpc
  cd grpc
  git submodule update --init
  sudo mkdir -p cmake/build
  cd cmake/build
  sudo cmake ../..
  sudo make protoc grpc_php_plugin
  sudo cp -rf grpc_php_plugin /usr/local/bin/grpc_php_plugin
  cd ${ROOT}
  sudo rm -rf grpc
  echo "grpc_php_plugin 安装完成"
else
  echo "grpc_php_plugin 已安装"
fi

echo "检查 protoc-gen-go..."
# protoc-gen-go 检查
if [ ! -f /usr/local/bin/protoc-gen-go ]; then
  echo "protoc-gen-go 未安装，正在安装..."
  sudo go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
  sudo cp -rf $HOME/go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go
  echo "protoc-gen-go 安装完成"
else
  echo "protoc-gen-go 已安装"
fi

echo "检查 protoc-gen-go-grpc..."
# protoc-gen-go-grpc 检查
if [ ! -f /usr/local/bin/protoc-gen-go-grpc ]; then
  echo "protoc-gen-go-grpc 未安装，正在安装..."
  sudo go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
  sudo cp -rf $HOME/go/bin/protoc-gen-go-grpc /usr/local/bin/protoc-gen-go-grpc
  echo "protoc-gen-go-grpc 安装完成"
else
  echo "protoc-gen-go-grpc 已安装"
fi

# PHP 输出目录检查
echo "检查 PHP 输出目录..."
if [ ! -d "${DIST_PHP}" ]; then
  echo "PHP 输出目录不存在，创建目录..."
  mkdir -p ${DIST_PHP}
else
  echo "PHP 输出目录已存在"
fi

# Go 输出目录检查
echo "检查 Go 输出目录..."
if [ ! -d "${DIST_GO}" ]; then
  echo "Go 输出目录不存在，创建目录..."
  mkdir -p ${DIST_GO}
else
  echo "Go 输出目录已存在"
fi

echo "开始生成 PB 文件..."
# 遍历 Proto 文件并生成代码
if [ -d "$PROTO_DIR" ]; then
  for file in "$PROTO_DIR"/*.proto; do
    if [ -f "$file" ]; then
      filename=$(basename "$file") # 提取文件名
      echo "处理文件: $filename"
      # 生成 PHP 文件
      echo "生成 PHP 文件..."
      protoc --proto_path=${PROTO_DIR} --php_out=${DIST_PHP} --grpc-php_out=${DIST_PHP} --plugin=protoc-gen-grpc-php=/usr/local/bin/grpc_php_plugin $filename
      echo "PHP 文件生成完成"
      # 生成 Go 文件
      echo "生成 Go 文件..."
      protoc --proto_path=${PROTO_DIR} --go_out=${DIST_GO} --go-grpc_out=${DIST_GO} $filename
      echo "Go 文件生成完成"
    fi
  done
else
  echo "Proto 文件目录 ${PROTO_DIR} 不存在！"
fi

echo "所有文件生成完成！"
