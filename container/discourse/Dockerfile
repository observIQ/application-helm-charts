FROM bitnami/discourse:2.8.9-debian-11-r3

RUN gem install prometheus_exporter

COPY plugin/discourse-prometheus /opt/bitnami/discourse/plugins/discourse-prometheus
RUN chown -R discourse:root /opt/bitnami/discourse/plugins/discourse-prometheus
