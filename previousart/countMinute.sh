#!/bin/bash

while read line
do
min=`echo $line -n| grep -Po '\[\K.*?:\d+:\d+'`;
 # ip=`echo -n $line|cut -d ' ' -f1`;
echo $min
done \
 | uniq -c 
# |sort -rn
