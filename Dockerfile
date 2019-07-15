# Multi-stage build
FROM golang:alpine AS builder
WORKDIR /websocket-echo-client
ADD . /websocket-echo-client
RUN apk update && apk add git && apk add ca-certificates
RUN cd /websocket-echo-client && go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o vue-websocket-echo

# Second stage, smaller image
FROM alpine
WORKDIR /websocket-echo-client
COPY --from=builder /websocket-echo-client/ /websocket-echo-client/
ENTRYPOINT ./vue-websocket-echo
