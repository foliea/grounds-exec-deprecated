#!/bin/sh

useradd $HOSTNAME

ulimit -p 15

echo "$1" > prog.c
gcc -o prog prog.c

if [ -f "prog" ]
then
  sudo -u $HOSTNAME ./prog
fi
