#!/bin/bash

echo "what is your name?"

read name

echo "How old are you?"

read age

echo "Calculating your future..."
sleep 3

echo "Hey, $name, you will become millionare in $((($RANDOM % 15)+ $age)) years."