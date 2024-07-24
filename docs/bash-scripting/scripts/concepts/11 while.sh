#!/bin/bash

x=1

while [[ $x -le 10 ]]
do
    echo "The number is $x"
#    x=$(( $x + 1 ))
    (( x++ ))
done
