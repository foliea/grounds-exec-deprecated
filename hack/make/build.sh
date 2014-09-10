#!/bin/sh

set -e

BUILD_CMD="docker build -t"

go() {
	${BUILD_CMD} "$GO_IMAGE" .
}

web() {
	${BUILD_CMD} "$WEB_IMAGE" web
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
