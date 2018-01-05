#!/bin/sh

CURRENT_DIR="$(cd $(dirname "$0") && pwd -P)"

cd "$CURRENT_DIR/rest" && \
go get -t -v && \
go build -o "$CURRENT_DIR/cryptorest"
