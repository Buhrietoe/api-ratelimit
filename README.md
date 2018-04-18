# api-ratelimit

Simple HTTP API server for demonstrating a proxied ratelimit.

## Usage:


## Build Options:

    ./build.sh

or

    docker build -t arle:latest .

## Run:

    ./arl -config arl.json

or

    docker run --rm -v arl.json:/arl.json arle:latest
