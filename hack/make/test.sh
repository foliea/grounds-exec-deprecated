#!/bin/sh

set -e

run="docker run --rm"

# Run binary compilation
binary() {
  $run "$GO_IMAGE" ./hack/binary.sh
}

# Run unit tests inside a docker container
unit() {
	$run "$GO_IMAGE" ./hack/test-unit.sh
}

# Run web tests inside a docker container
web() {
	$run -e "RAILS_ENV=test" "$WEB_IMAGE" rake test
}

test() {
	# If first parameter from CLI is missing or empty
	if [ -z $1 ]; then
		echo "usage: test [binary|unit|web]"
		return
	fi
	eval $1
}

test "$1"
