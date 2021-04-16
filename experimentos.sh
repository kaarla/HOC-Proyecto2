#!/bin/bash

for i in {1..1}
do
  seed=$(($(od -An -N3 -i /dev/random)))
  for j in 1
  do
    go run main.go problema $seed $j
  done
  echo "-------------------------------------"
done
