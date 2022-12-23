#!/bin/bash

echo "Hey, do you like coffee? (y/n)"

read answer

if [ $answer == "y" ]; then
    echo "Great, I like coffee too!"
else
    echo "Oh, I see. I like coffee."
fi