#!/bin/bash
i=0
while read line
do
	((i++));
	if [[ $i -eq 10 ]]; then
		echo $line
		exit 0
	fi
done < "./change.go"
