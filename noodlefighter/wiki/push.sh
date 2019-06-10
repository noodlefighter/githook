#!/bin/bash

echo "push.sh running!"

SHELL_DIR=$(cd "$(dirname "$0")";pwd)
cd $SHELL_DIR

cd build/wiki
git pull
./generate.sh
rsync -av ./hexo/public/ --delete /home/wwwroot/wiki




