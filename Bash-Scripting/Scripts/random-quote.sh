#!/bin/bash

read -p "How many quotes do you want to see? " NUM

if [[ -z $NUM ]]; then
    echo "Please enter a number"
    exit 1
elif [[ $NUM -eq 0 ]]; then
    echo "Please enter a number greater than 0"
    exit 1
elif [[ $NUM -lt 0 ]]; then
    echo "Please enter a positive number"
    exit 1
fi

URL=https://api.quotable.io/random

for (( i=1; i<=$NUM; i++ ))
do
    curl -s $URL | jq -r .content
    sleep 1
done
