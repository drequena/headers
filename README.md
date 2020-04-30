# Headers

A small app to test _API GATEWAYS_ such as [Kong](https://konghq.com/kong/) locally.

## About
This is a very simple app that I've created to test Kong in my personal studies.
All tutorials about _API GATEWAYS_ always use [httpbin](https://httpbin.org/) to test its features. I wanted to test locally.

## Use
Execute and make a request.
```bash
./headers
2020/05/01 20:08:43 Starting Headers on port: :9090

curl -q localhost:9090
{"STATUS:" "200", "HOST:" "little-italy"}
```
The request will be logged on terminal. All headers will be printed.
```bash
Accept - [*/*]
User-Agent - [curl/7.58.0]
Path: /
Method: GET
```
If you want to change the default behavior (HTTP STATUS CODE), do a GET request to path "**/set/[CODE]**"

```bash
./headers
2020/05/01 20:08:43 Starting Headers on port: :9090
2020/05/01 20:34:31 Changing / status code to: 500

curl -q localhost:9090/set/500
{"STATUS:" "OK"}

curl -v localhost:9090
* Rebuilt URL to: localhost:9090/
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET / HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.58.0
> Accept: */*
> 
< HTTP/1.1 500 Internal Server Error
< Date: Fri, 01 May 2020 23:37:25 GMT
< Content-Length: 42
< Content-Type: text/plain; charset=utf-8
< 
{"STATUS:" "500", "HOST:" "little-italy"}
```

## Config
You can customize the listen port by exporting a ENVVAR called PORT.
```bash
export PORT=8080
./headers
2020/05/01 20:08:43 Starting Headers on port: :8080
```
### Build
This is a very simples app, you can build it just by typing _go build_
Use _make_ to build and create docker image
```bash
make build  #compiled the code
make docker #create docker image
or
make all    # do both
```
