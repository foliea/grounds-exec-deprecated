#!/bin/sh

set -e

run="docker run -d --name"

redis() {
  $run groundsredis dockerfile/redis
}

websocket() {
  $run groundsock "$GO_IMAGE" -p "$WEBSOCKET_PORT":"$WEBSOCKET_PORT"./hack/run.sh -d -p ":$WEBSOCKET_PORT"
}

web() {
	$run groundsweb -p "$WEB_PORT":"$WEB_PORT" -e RAILS_PORT="$WEB_PORT" --link groundsredis:redis "$WEB_IMAGE" rake run
}

run() {
	# If first parameter from CLI is missing or empty
	if [ -z $1 ]; then
		echo "usage: run [redis|websocket|web]"
		return
	fi
	eval $1
}

run "$1"
