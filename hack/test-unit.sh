#!/bin/bash
set -e

if [ ! "$GOPATH" ]; then
	echo >&2 'error: missing GOPATH; please see http://golang.org/doc/code.html#GOPATH'
	exit 1
fi

echo "Testing runner..."
gom test -cover ./pkg/runner

echo "Testing handler..."
gom test -cover ./pkg/handler

echo "Testing utils..."
gom test -cover ./pkg/utils
