#!/bin/bash

while read line
do
# min=`echo $line -n| grep -Po '\[\K.*?:\d+:\d+'`;
hour=`echo $line -n| grep -Po '\[\K.*?:\d+:'`;
 # ip=`echo -n $line|cut -d ' ' -f1`;
echo $hour
done \
 | uniq -c 
# |sort -rn
