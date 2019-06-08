#!/bin/bash

mkdir -p ./build && cd build
git clone https://github.com/noodlefighter/wiki.git
cd wiki
chmod +x ./install.sh && ./install.sh
chmod +x ./generate.sh


