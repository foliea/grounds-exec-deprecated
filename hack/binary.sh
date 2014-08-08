#!/bin/bash
set -e

NAME="server"
BINARY="bin/$NAME"

echo "Getting dependencies..."

go get ./$NAME ./...

echo "Creating binary: $BINARY"

go build -o $BINARY ./$NAME

echo "Created binary: $BINARY"
