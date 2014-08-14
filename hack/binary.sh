#!/bin/bash
set -e

NAME="groundsock"
BINARY="bin/$NAME"

if [ ! "$GOPATH" ]; then
	echo >&2 'error: missing GOPATH; please see http://golang.org/doc/code.html#GOPATH'
	exit 1
fi

echo "Creating binary: $BINARY"

gom build -o $BINARY ./cmd/$NAME

echo "Created binary: $BINARY"
