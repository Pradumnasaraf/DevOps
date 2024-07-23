#!/bin/bash

until [[ $number -eq 10 ]]
do
    echo "The number is $number"
    read number
done
echo "You have entered 10"