APP_NAME=vue-websocket-echo
FRONTEND_DIR=websocket-echo-client

all: build-vue-app statik go docker 
build: frontend go-build
frontend: build-vue-app statik
go-build: go
container: docker
statik: statik

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

.PHONY: statik
statik:
	statik -src=./websocket-echo-client/dist
