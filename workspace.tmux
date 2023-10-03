#!/bin/bash

tmux-workspace "GoWas" "lib" -c -v "cd test && make run && make dev && zsh" "nvim && zsh" \
