#!/bin/sh

set -e

build="docker build -t"

go() {
  $build "$GO_IMAGE" .
}

web() {
	$build "$WEB_IMAGE" web
}

build() {
	# If first parameter from CLI is missing or empty
	if [ -z $1 ]; then
		echo "usage: build [go|web]"
		return
	fi
	eval $1
}

build "$1"
