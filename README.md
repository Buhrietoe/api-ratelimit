# api-ratelimit

Simple HTTP API server for demonstrating a proxied ratelimit.

## Usage:


## Build Options:

    ./build.sh

or

    docker build -t arl:latest .

## Run:

    ./arl -config arl.json

or

    docker run --rm -v $(pwd):/conf arl
