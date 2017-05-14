#!/bin/sh

if [[ -n "$1" ]]; then
	PORT=$1
else
	PORT=9001
fi

docker run -p $PORT:8080 -it word