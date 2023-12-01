#!/bin/bash

file="input-new11.txt"

count=0

while IFS= read -r line; do
    #echo "for $line --> with ${#line}"
    count=$(expr $count + 1)
    length=${#line}
    for ((i = 0; i < length; i++)); do
        char="${line:i:1}"
        #echo "Character at position $i: $char"
        if [[ $char =~ [0-9] ]];then
            echo -n "$char"
        fi
    done
    echo ""
done <$file 
