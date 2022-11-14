FROM docker.io/bitnami/discourse:2.8.9

COPY plugin/discourse-prometheus /opt/bitnami/discourse/plugins/discourse-prometheus
RUN chown -R discourse:root /opt/bitnami/discourse/plugins/discourse-prometheus
