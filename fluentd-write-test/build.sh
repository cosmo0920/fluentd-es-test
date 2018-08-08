#!/usr/bin/env bash

go build -o write-record
docker build -t fluentd-write-test -f Dockerfile .
