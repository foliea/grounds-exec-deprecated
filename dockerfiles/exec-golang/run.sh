#!/bin/sh

echo $1 > main.go
go run main.go 2> error.log
cat error.log | cut -d':' -f 2-4 >&2
