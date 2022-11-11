FROM ubuntu as stage

WORKDIR /
RUN apt-get update
RUN apt-get install -y git
RUN git clone \
    https://github.com/discourse/discourse-prometheus.git \
    discourse-prometheus

FROM bitnami/discourse:2.8.9-debian-11-r3

RUN gem install prometheus_exporter

COPY --from=stage /discourse-prometheus /bitnami/discourse/plugins/discourse-prometheus
RUN chown -R discourse:root /bitnami/discourse/plugins/discourse-prometheus
