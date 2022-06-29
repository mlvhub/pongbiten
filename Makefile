.PHONY: run
run:
	go run main.go

.PHONY: wasm
wasm:
	./scripts/wasm.sh

.PHONY: server
server: wasm
	docker stop pongbiten-server || true
	docker run -d \
    -v $(PWD)/wasm-build:/web \
    -p 8080:8080 \
	--name pongbiten-server \
	--rm \
    halverneus/static-file-server:latest
	open http://localhost:8080

.PHONY:deploy
deploy: wasm
	surge ./wasm-build https://pongbiten.surge.sh

.PHONY: android
android:
	ebitenmobile bind -target android -javapkg sh.surge.pongbiten -o mobile/android/pongbiten/pongbiten.aar ./mobile

.PHONY: install
install:
	gomobile install github.com/mlvhub/pongbiten