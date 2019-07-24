# Multi-stage build 

# Build production vue.js app
FROM node:lts-alpine AS frontend-builder
ARG VUE_APP_BASE_PATH
ENV VUE_APP_BASE_PATH $VUE_APP_BASE_PATH
WORKDIR /websocket-echo-client
ADD . /websocket-echo-client
RUN npm --prefix ui install
RUN npm run --prefix ui build

# Build go binary
FROM golang:alpine AS builder
WORKDIR /websocket-echo-client
COPY --from=frontend-builder /websocket-echo-client/ /websocket-echo-client/
RUN apk update && apk add git && apk add ca-certificates
RUN go get -d -v
RUN go get github.com/rakyll/statik
RUN statik -src=./ui/dist
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o vue-websocket-echo

# Copy final build to minimal container
FROM alpine
WORKDIR /websocket-echo-client
COPY --from=builder /websocket-echo-client/vue-websocket-echo /websocket-echo-client/vue-websocket-echo
ENTRYPOINT ./vue-websocket-echo
