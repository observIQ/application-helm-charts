FROM ubuntu:bionic

RUN apt-get update

# All of this is necessary to build/run mesos
RUN apt-get install -y \
wget \
curl \
binutils \
vim \
g++ \
openjdk-8-jdk \
patch \
libapr1-dev \
libcurl4-openssl-dev \
maven \
libsasl2-dev \
libsvn-dev \
libz-dev \
python-dev \
make \
autoconf \
libtool

# Download Mesos install files and Mesos Prometheus exporter
RUN wget https://dlcdn.apache.org/mesos/1.11.0/mesos-1.11.0.tar.gz \ 
&& wget https://github.com/mesos/mesos_exporter/releases/download/v1.1.2/mesos_exporter-1.1.2.linux-arm64.tar.gz \
&& tar -xvzf mesos-1.11.0.tar.gz \
&& tar -xzvf mesos_exporter-1.1.2.linux-arm64.tar.gz \
&& cd /mesos-1.11.0 \
# Set JAVA version to be used to OpenJDK 8
&& update-alternatives --set java /usr/lib/jvm/java-8-openjdk-arm64/jre/bin/java \
# Build/Install Mesos
&& ./configure \
&& make install \
# Remove unnecessary files to maintain a slimmer image size
&& cd / \
&& rm -rf mesos-1.11.0 \
&& rm mesos-1.11.0.tar.gz
