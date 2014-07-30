#!/bin/sh

echo "\$stdout.sync = true\n$1" > prog.rb
ruby prog.rb
