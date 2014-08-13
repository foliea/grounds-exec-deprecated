FROM ubuntu:14.04

RUN apt-get update -qq

RUN apt-get -qy install \
    mono-mcs

COPY run.sh csharp/run.sh

WORKDIR csharp

ENTRYPOINT ["/csharp/run.sh"]
