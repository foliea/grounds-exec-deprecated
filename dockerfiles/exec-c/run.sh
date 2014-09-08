#!/bin/sh

useradd $HOSTNAME
su $HOSTNAME
ulimit -p 15

echo "$1" > prog.c
gcc -o prog prog.c

if [ -f "prog" ]
then
  ./prog
fi
