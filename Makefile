
UPX:=$(shell which upx)
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')

all: http-proxy

.PHONY: http-proxy
http-proxy:
	go build -ldflags "-X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\"" -o ./bin/http-proxy *.go
	if test -x "${UPX}"; then ${UPX} ./bin/http-proxy; else echo "upx not found"; fi

.PHONY: clean
clean:
	rm -f ./bin/http-proxy
