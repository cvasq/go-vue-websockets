# go-vue-websockets

A simple Go web server hosting a Vue.JS web application

The Vue.js app initiates a websocket connection the WebSocket Echo Server at echo.websocket.org


**Running the app via Go**

1. Clone this repo:
```
git clone git@github.com:carlos-py/go-vue-websockets.git
```

2. Enter directory
```
cd go-vue-websockets
```

3. Run the Go app with an optional port
```
$ go run main.go --port 8000

2018/08/12 15:14:55 HTTP server started on port: 8000

```
**Running the app via Docker**

1. Clone and enter repo
2. Build docker image
Replace your username
```
docker build -t <username>/go-websockets-vue .
```
3. Run container and map listening port to localhost.
Listening port set in Dockerfile
```
docker run -it --rm -p 8000:8000 toor/go-websockets-vue:latest

Output:
2018/08/12 23:12:38 HTTP server started on port: 8000

```