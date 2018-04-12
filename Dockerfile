FROM golang:1.10.1 AS build
COPY . /go/src/github.com/Buhrietoe/api-ratelimit-example/
WORKDIR /go/src/github.com/Buhrietoe/api-ratelimit-example/
ENV CGO_ENABLED 0
RUN go build -v -ldflags "-s -w" -o arle .

FROM scratch
LABEL maintainer "Jason Gardner <buhrietoe@gmail.com>"
EXPOSE 8080
COPY --from=build /go/src/github.com/Buhrietoe/api-ratelimit-example/arle /arle
COPY --from=build /go/src/github.com/Buhrietoe/api-ratelimit-example/arle.json /arle.json
CMD ["/arle", "-config", "arle.json"]
