#!/bin/bash

# Declare a 2-dimensional array with initial values
declare -A my_array=(
  [0,0]="11.20.10.21"
  [0,1]="11.20.10.22"
  [1,0]="11.30.10.21"
  [1,1]="11.30.10.22"
)

# Access the values in the array
echo ${my_array[0,0]}  # Output: 1
echo ${my_array[1,0]}  # Output: 3

# Get the dimensions of the array
rows=2
cols=2

# Iterate over the values in the array
for ((i=0; i<rows; i++)); do
  for ((j=0; j<cols; j++)); do
    echo ${my_array[$i,$j]}
  done
done
