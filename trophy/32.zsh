#!/usr/bin/env zsh

# title: 歴史の書き換え
# desc: git push -fした\nこれまでのこと、全部なかったことにしよう
# grade: silver

command grep -e 'git push.*(-f|--force)' &>/dev/null
