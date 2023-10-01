#!/bin/bash

tmux-workspace "GoWas" "lib" -c "nvim && zsh" \
	-w "test" -c "cd test && nvim" "cd test && make dev" \
	-w "server" -c "cd test && make run &"
