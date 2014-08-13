FROM ubuntu:14.04

RUN apt-get update -qq

RUN apt-get -qy install \
    build-essential \
    gcc

COPY run.sh cpp/run.sh

WORKDIR cpp

ENTRYPOINT ["/cpp/run.sh"]
