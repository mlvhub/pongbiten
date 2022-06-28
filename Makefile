.PHONY: run
run:
	go run main.go

.PHONY: wasm
wasm:
	./scripts/wasm.sh

.PHONY: server
server: wasm
	docker stop pong-server || true
	docker run -d \
    -v $(PWD)/build:/web \
    -p 8080:8080 \
	--name pong-server \
	--rm \
    halverneus/static-file-server:latest
	open http://localhost:8080

.PHONY:deploy
deploy: wasm
	surge ./build https://pongbiten.surge.sh