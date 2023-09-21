#!/bin/zsh

ls ../**/*.(go|html|js) | entr make remake
