#!/bin/bash

SHELL_DIR=$(cd "$(dirname "$0")";pwd)
cd $SHELL_DIR

go build
./githook $1
