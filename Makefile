APP_NAME=vue-websocket-echo
FRONTEND_DIR=websocket-echo-client

all: build-vue-app build-go docker
frontend: build-vue-app
local-go: frontend build-go
container: frontend docker

.PHONY: build-vue-app
build-vue-app: 
	npm run --prefix $(FRONTEND_DIR) build

.PHONY: build-go
build-go:
	go build .

.PHONY: docker
docker:
	docker build -t cvasquez/$(APP_NAME) .
