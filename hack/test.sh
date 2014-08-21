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

# Run all tests inside docker containers
all() {
	unit
	web
}

# Run unit tests inside a docker container
unit() {
	docker run --rm $SERVER_IMAGE sh test-unit.sh
}

# Run web tests inside a docker container
web() {
	docker run --rm -e "RAILS_ENV=test" $WEB_IMAGE bundle exec rspec
}

test() {
	# If first parameter from CLI is missing or empty
	if [ -z $1 ]; then
		echo "usage: [all|unit|web]"
	return
	fi
	eval $1
}

test $1
