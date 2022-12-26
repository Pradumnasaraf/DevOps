#!/bin/bash

for x in google.com yahoo.com facebook.com google.coom
do
    if ping -c 1 $x > /dev/null
    then
        echo "$x is up"
    else
        echo "$x is down"
    fi
done
