FROM fluent/fluentd:v1.2.5-debian

RUN fluent-gem install fluent-plugin-elasticsearch

ENV FLUENTD_CONF=fluent.conf

COPY *.* /fluentd/etc/

