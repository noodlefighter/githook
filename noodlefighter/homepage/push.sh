#!/bin/bash

SHELL_DIR=$(cd "$(dirname "$0")";pwd)
cd $SHELL_DIR

cd build/homepage
git pull
./generate.sh
rsync -av ./hexo/public --delete /home/wwwroot/noodlefighter.com


