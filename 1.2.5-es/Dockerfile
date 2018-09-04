FROM fluent/fluentd:v1.2.5-debian

RUN apt-get update && apt-get install -y build-essential dh-autoreconf ruby-dev
RUN fluent-gem install typhoeus fluent-plugin-elasticsearch

ENV FLUENTD_CONF=fluent.conf

COPY *.* /fluentd/etc/
