#!/bin/bash

file="input-new1.txt"

count=0

exception=("2ne" "on8" "8wo")
front=("2ne" "1ight" "8wo")
last=("tw1" "on8" "eigh2")

while IFS= read -r line; do
    #echo -n "\nfor $line --> with ${#line}"
    #count=$(expr $count + 1)
    new_line=$line
    i=0
    
    length=${#line}
    #echo -n "$line --> "
    substr_index=0
    for substr in ${exception[@]}; do
        rest=${line#*$substr}
        strpos=$(( ${#line} - ${#rest} - ${#substr} ))
        if [[ $strpos -gt 0 ]]; then 
            echo "substring pos = $strpos, str = $line , substring = $substr"
            i=0
            first_char=0
            last_char=0
            for ((i = 0; i < length; i++)); do
                char="${line:i:1}"
                if [[ $char =~ [0-9] ]];then
                    if [[ first_char -eq 0 ]]
                    first_char=$i

                    if [[ strpos -le $i ]]; then
                        new_line=${new_line//$substr/${front[$substr_index]}}

                    fi

                fi
            done
        fi
        #echo -n " checking $number -> digit ${digits[$i]} before = $new_line"
        #new_line=${new_line//$number/${digits[$i]}}
        #echo -n " after = $new_line"
        #i=$(expr $i + 1)
    done
    #echo "$new_line"
done <$file 
