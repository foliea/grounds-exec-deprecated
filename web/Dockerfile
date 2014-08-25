FROM ubuntu:14.04

# Install dependencies.
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
    python-software-properties \
    nodejs \
    libqt4-dev \
    xvfb

# Install ruby-2.1.2.
RUN curl -O http://ftp.ruby-lang.org/pub/ruby/2.1/ruby-2.1.2.tar.gz && \
    tar -xzvf ruby-2.1.2.tar.gz && \
    cd ruby-2.1.2 && \
    ./configure --disable-install-doc && \
    make && \
    make install && \
    cd .. && \
    rm -r ruby-2.1.2 ruby-2.1.2.tar.gz && \
    echo 'gem: --no-document' > /usr/local/etc/gemrc

# Install bundler
RUN gem install bundler

# Copy the Gemfile and Gemfile.lock into the image.
COPY Gemfile rails/Gemfile
COPY Gemfile.lock rails/Gemfile.lock

# Install ruby gems.
RUN cd rails && bundle install

# Using xvfb to emulate a X server
# (required by capybara-webkit)
ENV DISPLAY localhost:1.0

# Everything up to here was cached. This includes
# the bundle install, unless the Gemfiles changed.

# Now copy the app into the image.
COPY . rails

# Set the final working dir to the Rails app's location.
WORKDIR rails




