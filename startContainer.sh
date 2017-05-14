#!/bin/bash

while getopts ":i:p:" opt; do
  case $opt in
    i) BUILD="$OPTARG"
    ;;
    p) PORT="$OPTARG"
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    ;;
  esac
done

if [[ $BUILD != "no" ]]; then
	docker build -t random_word_translator .
fi

if [[ -z $PORT ]]; then
	PORT=9001
fi

if docker run -d --name random_word_translator -p $PORT:8080 -it random_word_translator; then
	echo "The application should now be reachable at http://localhost:$PORT/word/"
else
	echo "There was an error starting docker. Please make sure any other containers are killed and removed by using:"
	echo "docker kill random_word_translator"
	echo "docker rm random_word_translator"
fi