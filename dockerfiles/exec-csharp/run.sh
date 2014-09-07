#!/bin/sh

ulimit -p 15

echo "$1" > prog.cs
mcs prog.cs

if [ -f "prog.exe" ]
then
  mono prog.exe
fi
