#!/usr/bin/env bash

set -ex

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/main ./src/main.go
