#!/bin/sh

ulimit -p 10

echo "$1" > prog.go
go run prog.go
