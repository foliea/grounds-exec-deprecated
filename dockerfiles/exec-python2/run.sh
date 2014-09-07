#!/bin/sh

ulimit -p 15

echo "$1" > prog.py
python2 prog.py
