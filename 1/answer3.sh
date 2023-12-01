#!/bin/bash

file="output2-new11.txt"

count=0

while IFS= read -r line; do
    echo "for $line starting count = $count"
    count=$(expr $count + $line)
    echo "for $line end count = $count"
    echo ""
done <$file 
echo "total  = $count"

