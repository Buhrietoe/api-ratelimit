# api-ratelimit

Simple Example HTTP API server for rate-limiting and proxying requests.


## Build/Run Locally:

    ./build.sh
    ./arl -config arl.json


## Build/Run in Docker:

    docker build -t arl:latest .
    docker run --rm -v $(pwd):/conf arl

The produced docker image is typically 4MB :)


## Configuration

There is an example configuration in arl.json.

The config file specifies the server ip and port to listen on, the remote ip and port a request should be proxied to, and a rate limit number in requests per second.


## About

api.go implements an example API call StoreEvent that accepts an HTTP POST with an event parameter. Calling this method is rate-limited.
