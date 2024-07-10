FROM ubuntu:jammy

# Add community repo and install Apache Mesos
RUN apt-get update
RUN apt-get install -y wget curl gnupg software-properties-common
RUN curl https://www.aventer.biz/assets/support_aventer.asc | apt-key add -
RUN add-apt-repository 'deb http://rpm.aventer.biz/Ubuntu/jammy jammy main'
RUN apt-get install -y aventer-mesos

# Download and extract the Apache Mesos Prometheus exporter
ENV EXPORTER_URL=""
RUN ARCH=$(uname -m) && \
    echo $ARCH && \
    if [ "${ARCH}" = "arm64" ] || [ "${ARCH}" = "aarch64" ]; then \
    EXPORTER_URL="https://github.com/mesos/mesos_exporter/releases/download/v1.1.2/mesos_exporter-1.1.2.linux-arm64.tar.gz"; \
    else \
    EXPORTER_URL="https://github.com/mesos/mesos_exporter/releases/download/v1.1.2/mesos_exporter-1.1.2.linux-amd64.tar.gz"; \
    fi && \
    wget ${EXPORTER_URL} && \
    tar -xzvf $(basename ${EXPORTER_URL})
