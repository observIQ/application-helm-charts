FROM alpine:3.18

# Build the Aerospike benchmark binary
RUN apk add go git && \
    git clone https://github.com/aerospike/aerospike-client-go.git && \
    cd aerospike-client-go/tools/benchmark && \
    go build . && \
    mv benchmark /etc && \
    rm -rf /aerospike-client-go && \
    apk del go git

ENTRYPOINT [ "/etc/benchmark" ]
