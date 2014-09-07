FROM google/golang:latest

COPY run.sh /gopath/run.sh

RUN useradd dev
RUN chown -R dev: /gopath

RUN echo "dev hard nproc 5" > /etc/security/limits.conf

WORKDIR /gopath
USER dev

ENTRYPOINT ["/gopath/run.sh"]
