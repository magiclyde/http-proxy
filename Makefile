
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')
UPX:=$(shell which upx)
PWD:=$(shell pwd)

all: http-proxy

.PHONY: http-proxy
http-proxy:
	go build -ldflags "-X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\"" -o ./bin/http-proxy *.go
	if test -x "${UPX}"; then ${UPX} ./bin/http-proxy; else echo "upx not found"; fi

.PHONY: install-upx
install-upx:
	sudo apt install upx-ucl

.PHONY: clean
clean:
	rm -f ./bin/http-proxy

.PHONY: docker-build
docker-build:
	docker build -t http-proxy:1.0 .

.PHONY: docker-run
docker-run:
	docker run -p 8888:8888 -v "${PWD}/etc:/app/etc" -v "${PWD}/assets/tls:/app/assets/tls" http-proxy:1.0