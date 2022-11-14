FROM ruby

RUN gem install prometheus_exporter

ENTRYPOINT [ "/usr/local/bundle/bin/prometheus_exporter", "--bind", "0.0.0.0", "--port",  "9394"]
