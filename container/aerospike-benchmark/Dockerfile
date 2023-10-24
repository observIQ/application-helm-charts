FROM golang:1.21.3-alpine3.18 as build
WORKDIR /
RUN apk add curl unzip
RUN \
    curl -sL \
    -o aerospike.zip \
    https://github.com/aerospike/aerospike-client-go/archive/refs/tags/v6.13.0.zip
WORKDIR /app
RUN unzip /aerospike.zip
WORKDIR /app/aerospike-client-go-6.13.0/tools/benchmark
RUN go build .
RUN mv benchmark /

FROM alpine:3.18
COPY --from=build /benchmark /benchmark
ENTRYPOINT [ "/benchmark" ]
