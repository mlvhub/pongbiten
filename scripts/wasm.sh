#!/usr/bin/env bash

rm -rf ./build
mkdir build
GOOS=js GOARCH=wasm go build -o ./build/pong.wasm pong
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build/
cp index.html ./build