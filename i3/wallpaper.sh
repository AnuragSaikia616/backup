#!/usr/bin/zsh
exec $(ls ~/Downloads/wallpapers/ | fzf --preview="wal -i ~/Downloads/wallpapers/{} --backend wal")
