# api-ratelimit

Simple Example HTTP API server for rate-limiting and proxying requests.


## Configuration

There is an example configuration in arl.json.

The config file specifies the server ip and port to listen on, the remote ip and port a request should be proxied to, and a rate limit in requests per second.

You can also pass configuration with the following environment variables:

    ARL_SERVER_HOST
    ARL_SERVER_PORT
    ARL_REMOTE_HOST
    ARL_REMOTE_PORT
    ARL_RATE_LIMIT


## About

api.go implements an example API call 'StoreEvent' that accepts an HTTP POST with an event parameter. Calling this method is rate-limited.

Example requests to show rate-limiting in action:

    # for i in {0..9}; do curl -sL 'http://127.0.0.1:8080/StoreEvent?event={"stuff":"things"}' -X POST; done


## Build/Run Locally:

    ./build.sh
    ./arl -config arl.json


## Build/Run in Docker:

    docker build -t arl:latest .
    docker run --rm -v $(pwd):/conf --network host arl

The produced docker image is typically 4MB :)
