#!/bin/sh

set -e

if [ -z $GIT_BRANCH ]; then
	GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null)
fi

# Set default docker repository if none exist in env
if [ -z $REPOSITORY ]; then
	REPOSITORY="grounds"
fi

# Set default groundsock image if none exist in env
if [ -z $GO_IMAGE ]; then
	GO_IMAGE="$REPOSITORY/go:$GIT_BRANCH"
fi

# Set default web image if none exist in env
if [ -z $WEB_IMAGE ]; then
	WEB_IMAGE="$REPOSITORY/web:$GIT_BRANCH"
fi

# Set default web port if none exist in env
if [ -z $WEB_PORT ]; then
	WEB_PORT="3000"
fi

# Set default web port if none exist in env
if [ -z $WEBSOCKET_PORT ]; then
	WEBSOCKET_PORT="8080"
fi

main() {
	# If first parameter from CLI is missing or empty
	if [ -z $1 ]; then
		echo "usage: make [build|test|images]"
		return
	fi
	REPOSITORY="$REPOSITORY" GO_IMAGE="$GO_IMAGE" WEB_IMAGE="$WEB_IMAGE" \
	WEB_PORT="$WEB_PORT" WEBSOCKET_PORT="$WEBSOCKET_PORT" \
	sh hack/make/"$1".sh "$2"
}

main "$@"
