#!/bin/bash
set -e

BINARY=bin/server

echo "Getting dependencies..."

go get ./server ./...

echo "Creating binary: $BINARY"

go build -o $BINARY ./server

echo "Created binary: $BINARY"
