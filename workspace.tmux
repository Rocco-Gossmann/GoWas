#!/bin/bash

tmux-workspace "GoWas" "editor" -c "nvim && zsh" -w "server" -c "cd test && make run && make dev && zsh"  \
