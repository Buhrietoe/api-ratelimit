#!/bin/bash

BIN_NAME=arle

echo Building...
CGO_ENABLED=0 go build -v -ldflags '-w -s' -o $BIN_NAME .
file $BIN_NAME
