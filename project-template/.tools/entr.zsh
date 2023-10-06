#!/bin/zsh

ls ../**/*.(go|html|js|png) | entr make remake
