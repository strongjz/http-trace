# Basic go commands
GOCMD=go
GOVER=1.12
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BUILDDIR=bin
VERSION="0.0.1"

# Binary names
BINARY_NAME=http-trace
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WIN=$(BINARY_NAME)_windows.exe

all: test build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		govendor sync

# Cross compilation
build-windows:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILDDIR)/$(BINARY_WIN) -v

build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILDDIR)/$(BINARY_UNIX) -v

docker-build:
		docker run --rm -it -v $(GOPATH):/go -w /go/src/github.com/strongjz/http-trace golang:$(GOVER) go build -o $(BINARY_UNIX) -v

docker-image:
		docker build -t strongjz/http-trace:$(VERSION) ./

docker-push: docker-image
		docker push strongjz/http-trace:$(VERSION)

build-all: build-windows build-linux