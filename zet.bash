#!/bin/bash

dt=$(date +%Y%m%d%H%M%S)
echo "name the new zet"
read name

mkdir $HOME/zets/${dt}-${name}
touch $HOME/zets/${dt}-${name}/README.md
vim $HOME/zets/${dt}-${name}/README.md
