FROM ubuntu:14.04

ENV RUBY_MAJOR 2.1
ENV RUBY_MINOR 2.1.2

RUN apt-get update -q && \
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

RUN curl -O http://ftp.ruby-lang.org/pub/ruby/$RUBY_MAJOR/ruby-$RUBY_MINOR.tar.gz && \
    tar -xzvf ruby-$RUBY_MINOR.tar.gz && \
    cd ruby-$RUBY_MINOR && \
    ./configure --disable-install-doc && \
    make && \
    make install && \
    cd .. && \
    rm -r ruby-$RUBY_MINOR ruby-$RUBY_MINOR.tar.gz && \
    echo 'gem: --no-document' > /usr/local/etc/gemrc

COPY run.sh /home/dev/run.sh

RUN useradd dev
RUN chown -R dev: /home/dev

WORKDIR /home/dev
USER dev

ENTRYPOINT ["/home/dev/run.sh"]
