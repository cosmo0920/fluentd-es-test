# fluentd-es-test

fluentd es plugin test

## ES service

```
docker run -it --rm --network=host -e ES_JAVA_OPTS="-Xms2g -Xmx2g" -e bootstrap_memory_lock=true -e discovery.type=single-node docker.elastic.co/elasticsearch/elasticsearch:5.4.3
```

## Kibana service

```
docker run -it --rm --network=host -e ELASTICSEARCH_URL=http://127.0.0.1:9200 docker.elastic.co/kibana/kibana:5.4.3
```

## Fluentd service(1.2.2-es)

1. build image
2. docker run with host network

## Write log(fluentd-write-test)

fluentd-write-test is a simple program, log to stdout, the command line parameter contral how many logs to write

```
docker run -d --log-driver fluentd --log-opt fluentd-async-connect --log-opt "fluentd-address=127.0.0.1:24224" --log-opt tag="docker.{{.ImageName}}" fluentd-write-test 1000000
```

## memory

```
docker stats $fluentd-container-id
docker top $fluentd-container-id
pmap -d -p $fluentd-plugin-pid
```