#!/bin/bash
exist=`docker ps -a -q -f=name=gink`
if [ -n "${exist}" ]; then
    docker rm -f gink > /dev/null
fi
cd ../ && docker build -t gink/gink:latest .
docker run --name=gink -d -p 80:80 gink/gink:latest