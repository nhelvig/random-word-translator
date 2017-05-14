#!/bin/sh

docker build -t random_word_translator .

if [[ -n "$1" ]]; then
	PORT=$1
else
	PORT=9001
fi
echo "The application should be reachable at http://localhost:$PORT/word/"

docker run -p $PORT:8080 -it random_word_translator