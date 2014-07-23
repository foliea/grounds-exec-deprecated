#!/bin/sh

echo $1 > main.go 
exec go run main.go 
