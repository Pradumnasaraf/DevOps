#!/bin/bash

read -p "Enter the IP address to ping: " ip

COUNTER=0

while true
do
    if ping -c 1 $ip > /dev/null; then
        echo "Hey, $ip is up!!"
        COUNTER=0
        break
    else
        COUNTER=$((COUNTER+1))
        echo "$COUNTER) Hey, $ip is down!!"
    fi
sleep 4
done