#!/bin/sh

set -e

# Run unit tests inside a docker container
unit() {
	docker run "$GO_IMAGE" ./hack/test-unit.sh
}

# Run web tests inside a docker container
web() {
	docker run -e "RAILS_ENV=test" "$WEB_IMAGE" rake test
}

test() {
	# If first parameter from CLI is missing or empty
	if [ -z $1 ]; then
		echo "usage: test [unit|web]"
		return
	fi
	eval $1
}

test "$1"
