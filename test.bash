#!/usr/bin/env bash

set -e

if [ ! -f "./bin/http-proxy" ]; then
    echo "missing http-proxy binary, build..."
    make
fi

# run http-proxy
LOGFILE=$(mktemp -t http-proxy.XXXXXXX)
echo "starting http-proxy"
echo "logging to $LOGFILE"
./bin/http-proxy >$LOGFILE 2>&1 &
PID=$!

sleep 2

cleanup() {
    echo "killing http-proxy PID $PID"
    kill -s TERM $PID || cat $LOGFILE
}
trap cleanup INT TERM EXIT

#curl -x http://localhost:8888 "https://httpbin.org/get"
curl -Lv -m 30 --proxy https://localhost:8888 --proxy-cacert ./assets/tls/server.pem "https://httpbin.org/get"
