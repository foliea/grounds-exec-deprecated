FROM ubuntu:14.04

RUN apt-get update -qq && \
    apt-get install -qy \
    python \
    python-pip \ 
    python-dev

COPY run.sh /home/dev/run.sh

RUN useradd dev
RUN chown -R dev: /home/dev

RUN echo "dev hard nproc 5" > /etc/security/limits.conf

WORKDIR /home/dev
USER dev

ENTRYPOINT ["/home/dev/run.sh"]
