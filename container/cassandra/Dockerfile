FROM cassandra:4.1.0

ARG rmi_addr

# Disable auth requirements
RUN sed -i 's/JVM_OPTS="$JVM_OPTS -Dcom.sun.management.jmxremote.authenticate=true"/JVM_OPTS="$JVM_OPTS -Dcom.sun.management.jmxremote.authenticate=false"/g' /etc/cassandra/cassandra-env.sh
RUN sed -i '/password/d' /etc/cassandra/cassandra-env.sh

ENV LOCAL_JMX="no"
ENV JVM_EXTRA_OPTS="-Djava.rmi.server.hostname=${rmi_addr}"
