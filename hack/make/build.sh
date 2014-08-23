#!/bin/sh

set -e

go() {
	docker build -t "$GO_IMAGE" .
}

web() {
	docker build -t "$WEB_IMAGE" web
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
