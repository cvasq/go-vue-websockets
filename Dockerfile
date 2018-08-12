FROM golang:alpine
WORKDIR /go-websockets-vue
ADD . /go-websockets-vue
RUN apk update && apk add git && apk add ca-certificates
RUN cd /go-websockets-vue && go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o go-websockets-vue

ENTRYPOINT ["./go-websockets-vue"]
