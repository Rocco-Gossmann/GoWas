#!/bin/zsh

kill `lsof -i :7353 | grep -w "(LISTEN)" | awk '{print $2}'`

