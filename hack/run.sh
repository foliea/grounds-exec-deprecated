#!/bin/bash
set -e

sh hack/binary.sh
bin/groundsock $1
