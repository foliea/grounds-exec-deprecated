#!/bin/bash
set -e

sh hack/binary.sh
bin/server $1
