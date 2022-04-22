#!/bin/bash

dt=$(date +%Y%m%d%H%M%S)

echo "new|existing dir"
read dir

if [[ ! -d $HOME/zet/zets/${dir} ]]
then
	mkdir $HOME/zet/zets/${dir}
fi

echo "name your zet"
read name

mkdir $HOME/zet/zets/${dir}/${name}-${dt}
touch $HOME/zet/zets/${dir}/${name}-${dt}/README.md

echo "vim into your new zet?"
read yn

if [[ yn != 'n' ]]
then
	vim $HOME/zet/zets/${dir}/${name}-${dt}/README.md
fi
