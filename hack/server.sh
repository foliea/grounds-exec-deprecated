#!/bin/bash
set -e

DEST=$1
BINARY=$DEST/server

go build -o $BINARY ./server

echo "Created binary: $BINARY"
