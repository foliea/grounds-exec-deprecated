#!/bin/bash
set -e

echo "Getting dependencies..."
go get ./execcode ./...

echo "Testing execcode"
go test ./execcode

echo "Getting dependencies..."
go get ./utils ./...

echo "Testing utils"
go test ./utils
