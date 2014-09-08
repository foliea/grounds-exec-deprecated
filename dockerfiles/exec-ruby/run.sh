#!/bin/sh

useradd $HOSTNAME
su $HOSTNAME
ulimit -p 15

echo "\$stdout.sync = true\n\$stderr.sync = true\n$1" > prog.rb
ruby prog.rb
