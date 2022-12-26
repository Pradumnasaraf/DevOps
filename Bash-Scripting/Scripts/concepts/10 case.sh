#!/bin/bash

echo "Pick a number from 1 to 3"
read number

case $number in
    1)
        echo "You picked 1"
        name="One"
        ;;
    2)
        echo "You picked 2"
        name="Two"
        ;;
    3)
        echo "You picked 3"
        name="Three"
        ;;
    *)
        echo "You did not pick a number from 1 to 3"
        ;;
esac

echo "$name"

