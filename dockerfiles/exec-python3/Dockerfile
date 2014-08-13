FROM ubuntu:14.04

RUN apt-get update -qq

RUN apt-get install -qy \
    python \
    python-pip \ 
    python-dev

COPY run.sh python/run.sh

WORKDIR python

ENTRYPOINT ["/python/run.sh"]
