# go-vue-websockets

A Golang HTTP server / Vue.js web application that establishes a Websocket connection to echo.websocket.org  

The frontend utilizes Vue.js Components and the entire application can be easily shipped as a binary or container thanks to Go + Docker :whale2:

Example project for starting other Vue + Go apps

**Directory structure**
```bash
├── go-vue-websockets
│   ├── ui - Vue.js frontend directory
│   └── statik - Embedded static frontend files directory
├── main.go - Main server entrypoint
├── server.go - HTTP Server logic
├── Makefile - Build helper
├── Dockerfile - Container build spec
```

**Build info and requirements:**

A Makefile is included for common build tasks.  
Built using Go 1.12 and Vue CLI 3

Build tools used:
- Vue CLI 
- NPM
- Golang 1.1+
- [Statik](https://github.com/rakyll/statik)
- Make command
- Docker

**List of routes**

URL Path | Request Type |Purpose
:-----:|:-----: |:-----:
`/`|GET|Main frontend
`/log-collector`|POST|Collects user entered input from the frontend
`/metrics`|GET|Prometheus style metrics endpoint

**Building with Docker**

1. Clone and enter this repo:
```
$ git clone git@github.com:cvasq/go-vue-websockets.git
```

2. Build the docker image (Container size: ~12MB)

```
$ make container
```
3. Run the built container image and map the application listening port to your localhost.  
_Default port: 8080_
```
$ docker run -p 8080:8080 cvasquez/vue-websocket-echo:latest  

Output:
2019/07/15 21:54:05 Server starting on port 8080... 
2019/07/15 21:54:05 Web Interface: http://localhost:8080/
2019/07/15 21:54:05 Prometheus Metrics: http://localhost:8080/metrics

```

**Building from source**

1. Clone and enter this repo:
```
$ git clone git@github.com:cvasq/go-vue-websockets.git
```
2. Build the Vue.JS frontend along with Go Binary
```
$ make go-build
```
3. Run the new executable binary file

```
$ ./go-vue-websockets 

Output:
2019/07/15 21:52:05 Server starting on port 8080... 
2019/07/15 21:52:05 Web Interface: http://localhost:8080/
2019/07/15 21:52:05 Prometheus Metrics: http://localhost:8080/metrics

```

**Help Message**
```
$ ./go-vue-websockets -h

NAME:
   go-vue-websockets - Vue.js Websocket Client Example

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
