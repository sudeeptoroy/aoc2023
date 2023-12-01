#!/bin/bash

file="input.txt"

count=0

numbers=("zero" "one" "two" "three" "four" "five" "six" "seven" "eight" "nine")
numbers=("nine" "eight" "seven" "six" "five" "four" "three" "two" "one" "zero")
digits=("9" "8" "7" "6" "5" "4" "3" "2" "1")

while IFS= read -r line; do
    #echo -n "\nfor $line --> with ${#line}"
    count=$(expr $count + 1)
    new_line=$line
    i=0
    #echo -n "$line --> "
    for number in ${numbers[@]}; do
        #echo -n " checking $number -> digit ${digits[$i]} before = $new_line"
        new_line=${new_line//$number/${digits[$i]}}
        #echo -n " after = $new_line"
        i=$(expr $i + 1)
    done
    echo "$new_line"
done <$file 
