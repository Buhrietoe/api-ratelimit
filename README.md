# api-ratelimit

Simple HTTP API server for demonstrating a proxied ratelimit.


## Build/Run Locally:

    ./build.sh
    ./arl -config arl.json


## Build/Run in Docker:

    docker build -t arl:latest .
    docker run --rm -v $(pwd):/conf arl
