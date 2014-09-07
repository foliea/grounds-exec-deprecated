#!/bin/sh

ulimit -p 5

echo "$1" > prog.py
python2 prog.py
