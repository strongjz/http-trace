FROM golang:1.12 AS builder

MAINTAINER strongjz

COPY ./ /go/src/github.com/strongjz/http-trace

WORKDIR /go/src/github.com/strongjz/http-trace

RUN make build-linux

FROM busybox

# Retrieve the binary from the previous stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/strongjz/http-trace/bin/http-trace_unix /usr/local/bin/http-trace

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "/usr/local/bin/http-trace"]