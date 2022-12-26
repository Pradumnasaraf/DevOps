#!/bin/bash

while true
do
    read -p "Enter a number: " num
    if [[ $num -eq 10 ]]; then
        break
    elif [[ $num -eq 5 ]]; then
        continue
    fi
    echo "You have entered $num"
done