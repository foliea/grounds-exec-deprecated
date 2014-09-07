#!/bin/sh

ulimit -p 5

echo "$1" > prog.go
go run prog.go
