# Dev Environment

To hack on `Grounds`, all you need is `git`, `make`, `docker` and your favorite text editor.

You can run tests and even the whole stack inside `docker` containers, with the same environment
used in production.

If you prefer to hack on your local environment, there is also instructions to install all
required dependencies.

## Requirements

You need first to install `docker`.

Checkout the official `docker` [documentation](https://docs.docker.com/installation/)
to install it on your platform.

## Inside docker containers

### Websocket Application

To build the websocket `docker` image:

    $ make build go

To run unit tests inside a `docker` container:

    $ make test-unit

### Web Application

To build the web `docker` image:

    $ make build web

To run tests inside a `docker` container:

    $ make test-web

## Locally

### Websocket Application

First, you need go-1.3.1 and a properly `GOPATH`.

This project use the Go package manager.

Install Go package manager:

    $ go get github.com/mattn/gom

Install go package dependencies:

    $ $GOPATH/bin/gom install

#### Compile the websocket server

    $ make binary

#### Run unit tests

    $ ./hack/test-unit.sh

### Web Application

First, you need ruby-2.1.2, and `bundle`.

Install ruby gem dependencies:

    $ bundle install

#### Run the web server:

Set required envirnment variables:

    $ export RUN_ENDPOINT="{url websocket server}/run"
    $ export REDIS_URL="{url redis}"

Then:

    $ bundle exec rails s

#### Run tests

    $ bundle exec rspec
