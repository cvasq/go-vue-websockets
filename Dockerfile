# Multi-stage build 

# Build production vue.js app
FROM node:lts-alpine AS frontend-builder
WORKDIR /websocket-echo-client
ADD . /websocket-echo-client
RUN npm --prefix websocket-echo-client install
RUN npm run --prefix websocket-echo-client build

# Build go binary
FROM golang:alpine AS builder
WORKDIR /websocket-echo-client
COPY --from=frontend-builder /websocket-echo-client/ /websocket-echo-client/
RUN apk update && apk add git && apk add ca-certificates
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o vue-websocket-echo

# Copy final build to minimal container
FROM alpine
WORKDIR /websocket-echo-client
COPY --from=builder /websocket-echo-client/ /websocket-echo-client/
ENTRYPOINT ./vue-websocket-echo
