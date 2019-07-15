# go-vue-websockets

A basic Vue.JS web application which inititates a Websocket connection to echo.websocket.org  
It uses Vue components on the frontend and the application can be easily shipped with Go and Docker

Example project for building other Vue + Go apps

A Makefile is included for basic build tasks.  
Built using Go 1.12

**Building with Docker**

1. Clone and enter this repo:
```
$ git clone git@github.com:cvasq/go-vue-websockets.git
```

2. Build the docker image
```
$ make container
```
3. Run the built container and map the application listening port to your localhost.  
_Default port 8080_
```
$ docker run -p 8080:8080 cvasquez/vue-websocket-echo:latest  

Output:
2019/07/15 22:16:58 Server listening on port 8080
2019/07/15 22:16:58 Access the web UI at http://localhost:8080/

```

**Building from source**

1. Clone and enter this repo:
```
$ git clone git@github.com:cvasq/go-vue-websockets.git

$ cd go-vue-websockets
```
2. Build the Vue.JS frontend and Go Binary
```
$ make local-go
```
3. Run the compiled binary

```
$ ./go-vue-websockets 
2019/07/15 18:23:00 Server listening on port 8080
2019/07/15 18:23:00 Access the web UI at http://localhost:8080/  

```

