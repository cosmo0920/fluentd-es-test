FROM fluent/fluentd:1.2.2-debian

RUN fluent-gem install fluent-plugin-elasticsearch

ENV FLUENTD_CONF=fluent.conf

COPY *.* /fluentd/etc/

