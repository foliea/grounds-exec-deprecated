#!/bin/sh

set -e

GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null)

# Set default docker repository if none exist in env
if [ -z $REPOSITORY ]; then
  REPOSITORY="grounds"
fi

# Set default groundsock image if none exist in env
if [ -z $SERVER_IMAGE ]; then
  SERVER_IMAGE="$REPOSITORY/groundsock:$GIT_BRANCH"
fi

# Set default web image if none exist in env
if [ -z $WEB_IMAGE ]; then
  WEB_IMAGE="$REPOSITORY/web:$GIT_BRANCH"
fi

server() {
	docker build -t $(SERVER_IMAGE) .
}

web() {
	docker build -t $(WEB_IMAGE) web
}

build() {
	# If first parameter from CLI is missing or empty
  if [ -z $1 ]; then
    echo "usage: [server|web]"
    return
  fi
  eval $1
}

build $1