#!/bin/sh

cd ./rest && \
go get -t -v && \
go build -o ./cryptorest
