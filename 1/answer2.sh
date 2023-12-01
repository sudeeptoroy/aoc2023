#!/bin/bash

file="output1-new11.txt"

count=0
first=0
last=0

while IFS= read -r line; do
    #echo -n "for $line"
    length=${#line}
    for ((i = 0; i < length; i++)); do
        char="${line:i:1}"
        if [[ i -eq 0 ]]; then
            first=$char
        fi
        last=$char
    done
    echo "$first$last"
done <$file 
