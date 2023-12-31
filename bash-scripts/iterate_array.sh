#!/bin/bash

# Declare an array
my_array=("apple" "banana" "cherry" "date")

# Iterate over the elements in the array
for element in "${my_array[@]}"
do
    echo "$element"
done

