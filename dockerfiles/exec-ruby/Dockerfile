FROM ubuntu:latest

RUN apt-get update

RUN apt-get -y install \
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

ADD run.sh ruby/run.sh

WORKDIR ruby

ENTRYPOINT ["/ruby/run.sh"]
