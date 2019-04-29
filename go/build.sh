#!/bin/bash

mkdir -p builds

rice embed-go

env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o builds/enjoy_darwin_amd64

env GOOS=linux GOARCH=amd64 go build -o builds/enjoy_linux_amd64

set CGO_ENABLED=1 
set GOOS=windows 
set GOARCH=amd64 
go build -o builds/enjoy_win_amd64
