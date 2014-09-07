#!/bin/sh

ulimit -p 15

echo "$1" > prog.go
go run prog.go
