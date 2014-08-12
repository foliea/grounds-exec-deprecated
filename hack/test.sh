#!/bin/bash
set -e

if [ ! "$GOPATH" ]; then
	echo >&2 'error: missing GOPATH; please see http://golang.org/doc/code.html#GOPATH'
	exit 1
fi

echo "Getting dependencies..."
go get ./execcode ./...

echo "Testing execcode"
go test ./execcode

echo "Getting dependencies..."
go get ./utils ./...

echo "Testing utils"
go test ./utils
