name := pongbiten
server_name := ${name}-server

.PHONY: run
run:
	go run main.go

.PHONY: wasm
wasm:
	./scripts/wasm.sh $(name)

.PHONY: server
server: wasm stop
	docker run -d \
    -v $(PWD)/wasm-build:/web \
    -p 8080:8080 \
	--name $(server_name) \
	--rm \
    halverneus/static-file-server:latest
	open http://localhost:8080?eruda=true

.PHONY: stop
stop:
	docker stop $(server_name) || true

.PHONY:deploy
deploy: wasm
	surge ./wasm-build https://$(name)2.surge.sh

.PHONY: android
android:
	ebitenmobile bind -target android -javapkg sh.surge.$(name) -o mobile/android/$(name)/$(name).aar ./mobile

.PHONY: install
install:
	gomobile install github.com/mlvhub/$(name)