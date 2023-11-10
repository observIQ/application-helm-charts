FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y wget unzip && \
    wget https://github.com/influxdata/influxdb2-sample-data/archive/refs/heads/master.zip && \
    unzip master.zip && \
    wget https://dl.influxdata.com/influxdb/releases/influxdb2-client-2.7.3-linux-amd64.tar.gz && \
    tar -xvf influxdb2-client-2.7.3-linux-amd64.tar.gz && \
    mv influx /usr/local/bin

WORKDIR /
COPY loadgen.sh /
RUN chmod +x /loadgen.sh

ENTRYPOINT [ "/loadgen.sh" ]

