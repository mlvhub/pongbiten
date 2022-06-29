#!/usr/bin/env bash

rm -rf ./wasm-build
mkdir wasm-build
GOOS=js GOARCH=wasm go build -o ./wasm-build/pongbiten.wasm github.com/mlvhub/pongbiten
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./wasm-build/
cp index.html ./wasm-build