FROM ubuntu:14.04

RUN apt-get update -qq && \
    apt-get install -qy \
    python \
    python-pip \ 
    python-dev

COPY run.sh /home/dev/run.sh

WORKDIR /home/dev

ENTRYPOINT ["/home/dev/run.sh"]
