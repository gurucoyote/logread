#!/bin/bash

while read line
do
  echo -n "$line " 
ip=`echo -n $line|grep -o '[^ ]*$'`
      dig -x "$ip" +short
done
