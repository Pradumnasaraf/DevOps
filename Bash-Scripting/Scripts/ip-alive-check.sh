#!/bin/bash

read -p "Enter the IP address to ping: " ip

while true
do
    if ping -c 1 $ip > /dev/null; then
        echo "Hey, $ip is up!!"
        break
    else
        echo "$ip is currently down"
    fi
sleep 2
done