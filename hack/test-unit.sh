#!/bin/sh

set -e

if [ ! "$GOPATH" ]; then
	echo >&2 'error: missing GOPATH; please see http://golang.org/doc/code.html#GOPATH'
	exit 1
fi

get_pkg_dirs() {
	echo $(find pkg -maxdepth 1 -type d | grep pkg/)
}

# For every pkg
for dir in $(get_pkg_dirs); do
	echo "Testing: $dir"
	gom test -cover "./$dir"
done

