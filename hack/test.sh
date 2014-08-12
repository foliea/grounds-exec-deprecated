#!/bin/bash
set -e

if [ ! "$GOPATH" ]; then
	echo >&2 'error: missing GOPATH; please see http://golang.org/doc/code.html#GOPATH'
	exit 1
fi

echo "Testing execcode"
gom test ./execcode

echo "Testing utils"
gom test ./utils
