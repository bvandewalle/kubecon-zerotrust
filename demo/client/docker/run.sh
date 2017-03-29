#!/bin/bash

while :
do
  curl -s --connect-timeout 1 --retry 0 $BACKEND
	sleep 1
done
