#!/bin/zsh

go run ./server/server.go & 
echo started http://localhost:7353
#killall `lsof -i :7353 | grep -w "(LISTEN)" | awk '{print $2}'`
