FROM cassandra:4.1.0

# Copy ready-probe script into container
COPY ./ready-probe.sh ./config.yaml /
RUN chmod +x /ready-probe.sh

# Download wget and jmx_exporter Java agent
RUN apt-get update && \
    apt-get install -y wget && \
    wget https://repo.maven.apache.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.20.0/jmx_prometheus_javaagent-0.20.0.jar

# Disable auth requirements
RUN sed -i 's/JVM_OPTS="$JVM_OPTS -Dcom.sun.management.jmxremote.authenticate=true"/JVM_OPTS="$JVM_OPTS -Dcom.sun.management.jmxremote.authenticate=false"/g' /etc/cassandra/cassandra-env.sh
RUN sed -i '/password/d' /etc/cassandra/cassandra-env.sh

ENV LOCAL_JMX="no"
ENV JVM_EXTRA_OPTS="-Djava.rmi.server.hostname=127.0.0.1 -javaagent:./jmx_prometheus_javaagent-0.20.0.jar=12345:/config.yaml"
