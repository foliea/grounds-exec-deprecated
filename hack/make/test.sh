#!/bin/sh

set -e

# Run all tests inside docker containers
all() {
	unit
	web
}

# Run unit tests inside a docker container
unit() {
	docker run --rm "$GO_IMAGE" ./test-unit.sh
}

# Run web tests inside a docker container
web() {
	docker run --rm -e "RAILS_ENV=test" "$GO_IMAGE" bundle exec rspec
}

test() {
	# If first parameter from CLI is missing or empty
	if [ -z $1 ]; then
		echo "usage: test [all|unit|web]"
		return
	fi
	eval $1
}

test "$1"
