FROM quay.io/wildfly/wildfly:27.0.0.Final-jdk17
RUN /opt/jboss/wildfly/bin/add-user.sh admin password --silent
COPY helloworld.war /opt/jboss/wildfly/standalone/deployments
COPY standalone.xml /opt/jboss/wildfly/standalone/configuration/standalone.xml
