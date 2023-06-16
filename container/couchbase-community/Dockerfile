FROM couchbase:community-7.2.0

# Config script
COPY ./base-config-script.sh /base-config-script.sh
RUN chmod +x /config-script.sh

# Start script
COPY ./start-configure-node.sh /start-configure-node.sh
RUN chmod +x /start-configure-node.sh

ENTRYPOINT [ "/bin/sh", "-c", "/start-configure-node.sh" ]
