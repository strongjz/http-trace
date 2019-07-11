# http-trace

[![Go Report Card](https://goreportcard.com/badge/github.com/strongjz/http-trace)](https://goreportcard.com/report/github.com/strongjz/http-trace)

Record information about HTTP requests

[Docker Repo](https://hub.docker.com/r/strongjz/http-trace)

Building both windows and linux 

```bash
make build-all 
```

Stand alone binary

```bash
./http-trace --method get --url "https://google.com/"
```

Create the docker image 
```bash
 make docker-image VERSION=0.0.2
```

Run http-tracer

```bash
docker run --rm -it strongjz/http-trace:0.0.2 --method get --url "https://google.com/"

2019/07/11 20:17:34 req: get / HTTP/1.1
Host: google.com

2019/07/11 20:17:34 DNS result google.com.      299     IN      A       172.217.6.110
2019/07/11 20:17:34 URL Scheme:         https
2019/07/11 20:17:34 URL host:   google.com
2019/07/11 20:17:34 URL Path:   /
2019/07/11 20:17:34 Name Server:        0xc00007ad00
2019/07/11 20:17:34 Connection Time:    125.480783ms
2019/07/11 20:17:34 DNS Lookup:         90.651642ms
2019/07/11 20:17:34 Name lookup:        90.651642ms
2019/07/11 20:17:34 Pretransfer:        261.766508ms
2019/07/11 20:17:34 Server Processing:  46.996239ms
2019/07/11 20:17:34 Start Transfer:     309.844351ms
2019/07/11 20:17:34 TCP Connection:     34.819129ms
2019/07/11 20:17:34 TLS  Handshake:     136.20805ms
2019/07/11 20:17:34 Status Code:        405
2019/07/11 20:17:34 Entire timing:      374.708861ms
```


