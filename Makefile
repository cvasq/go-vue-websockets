APP_NAME=vue-websocket-echo
FRONTEND_DIR=websocket-echo-client

all: build-vue-app go docker
frontend: build-vue-app
go-build: build-vue-app go
container: docker

.PHONY: build-vue-app
build-vue-app: 
	npm --prefix $(FRONTEND_DIR) install
	npm run --prefix $(FRONTEND_DIR) build

.PHONY: go
go:
	go build .

.PHONY: docker
docker:
	docker build -t cvasquez/$(APP_NAME) .
