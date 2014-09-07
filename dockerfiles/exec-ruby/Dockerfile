FROM ubuntu:14.04

RUN apt-get update -qq && \
    apt-get -qy install \
    git-core \
    curl \
    zlib1g-dev \
    build-essential \
    libssl-dev \
    libreadline-dev \
    libyaml-dev \
    libsqlite3-dev \
    sqlite3 \
    libxml2-dev \
    libxslt1-dev \
    libcurl4-openssl-dev \
    python-software-properties

RUN curl -O http://ftp.ruby-lang.org/pub/ruby/2.1/ruby-2.1.2.tar.gz && \
    tar -xzvf ruby-2.1.2.tar.gz && \
    cd ruby-2.1.2 && \
    ./configure --disable-install-doc && \
    make && \
    make install && \
    cd .. && \
    rm -r ruby-2.1.2 ruby-2.1.2.tar.gz && \
    echo 'gem: --no-document' > /usr/local/etc/gemrc

COPY run.sh /home/dev/run.sh

RUN useradd dev
RUN chown -R dev: /home/dev

RUN echo "dev hard nproc 5" > /etc/security/limits.conf

WORKDIR /home/dev
USER dev

ENTRYPOINT ["/home/dev/run.sh"]
