#!/usr/bin/env bash

# exit when any command fails
set -e

NAME=$1

mkdir -p wasm-build
rm -rf ./wasm-build/*
GOOS=js GOARCH=wasm go build -o ./wasm-build/game.wasm github.com/mlvhub/$NAME
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./wasm-build/
cp index.html ./wasm-build
