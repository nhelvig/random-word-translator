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

docker run -d --name random_word_translator -p $PORT:8080 -it random_word_translator 

echo "The application should now be reachable at http://localhost:$PORT/word/"