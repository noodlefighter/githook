#!/bin/bash

mkdir -p ./build && cd build
git clone https://github.com/noodlefighter/homepage.git
cd homepage
chmod +x ./install.sh && ./install.sh


