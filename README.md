# go-vue-websockets

A basic Vue.js web application which inititates a Websocket connection to echo.websocket.org  

The frontend utilizes Vue.js Components and the application can be easily shipped as a binary or container thanks to Go + Docker :whale2:

Example project for starting other Vue + Go apps


**Build info and requirements:**

A Makefile is included for basic build tasks.  
Built using Go 1.12 and Vue CLI 3

Build tools:
- NPM + Vue CLI
- Golang 1.2+
- Docker (optional)


**Building with Docker**

1. Clone and enter this repo:
```
$ git clone git@github.com:cvasq/go-vue-websockets.git
```

2. Build the docker image
```
$ make container
```
3. Run the built container image and map the application listening port to your localhost.  
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
```
2. Build the Vue.JS frontend and Go Binary
```
$ make local-go
```
3. Run the compiled binary

```
$ ./go-vue-websockets 

Output:
2019/07/15 18:23:00 Server listening on port 8080
2019/07/15 18:23:00 Access the web UI at http://localhost:8080/  

```

**Help Message**
```
$ ./go-vue-websockets -h

NAME:
   go-vue-websockets - Vue.js Websocket Example

USAGE:
   go-vue-websockets [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Carlos Vasquez <carlos@email.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --listening-port value  Listening Port (default: "8080") [$LISTENING_PORT]
   --help, -h              show help
   --version, -v           print the version
```
