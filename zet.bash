#!/bin/bash

dt=$(date +%Y%m%d%H%M%S)

echo "name the new zet"
read name

mkdir $HOME/zet/zets/${dt}-${name}
touch $HOME/zet/zets/${dt}-${name}/README.md
vim $HOME/zet/zets/${dt}-${name}/README.md
