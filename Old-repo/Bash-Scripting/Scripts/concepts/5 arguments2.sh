#!/bin/bash

name=$1

user=$(whoami)
date=$(date +%F)
whereami=$(pwd)

echo "Good Morning $name"

echo "You are curently logged in as $user and you are in the $whereami directory on $date"