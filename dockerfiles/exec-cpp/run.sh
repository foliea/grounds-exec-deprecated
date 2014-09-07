#!/bin/sh

ulimit -p 5

echo "$1" > prog.cpp
g++ -o prog prog.cpp

if [ -f "prog" ]
then
  ./prog
fi
