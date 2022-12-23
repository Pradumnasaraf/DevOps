#!/bin/bash

beast=$(( $RANDOM % 2 ))

echo "Your first besat approcahes. Prepare to fight!. Pick your weapon: (0)Sword (1)Gun"

read weapon

if [ $weapon == $beast ]; then
    echo "You have slain the beast!"
else
    echo "You have been slain by the beast!"
    exit 1
fi

sleep 2

echo "The fight is not over. You have been transported to a new location. Pick a number between 0 and 9"

read number

beast=$(( $RANDOM % 10 ))

if [ $number == $beast -o $number == "cheat" ]; then
    echo "You have slain the beast!"
else
    echo "You have been slain by the beast!"
    exit 1
fi