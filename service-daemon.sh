#!/bin/sh
fileName="goweb.service"

path=$(pwd)
pathFile=$path/background/$fileName
pathLinkFile=/etc/systemd/system/$fileName

if [ ! -e pathLinkFile ]; then
  echo "linked file background/goweb.service"
  ln -sf $pathFile $pathLinkFile
fi
