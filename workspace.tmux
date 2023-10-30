#!/bin/bash

tmux-workspace "GoWas" "editor" -c "cd test && make dev" -c "cd test && make run"
