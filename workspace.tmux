#!/bin/bash

tmux-workspace "GoWas" "editor" -c "nvim && zsh" "cd test && make dev" \
	-w "server" -c "cd test && make run"
