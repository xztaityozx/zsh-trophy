#!/usr/bin/env zsh

# title: 100回目のls
# desc: lsコマンドを100回実行した\nここまで来たらslコマンドとは間違わないよね？
# grade: gold

command grep -e '^ls' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="ls,count"
  if [[ "$ZT_RECORD_HASH[$k]" == "99" ]]; then
    true
  else
    false
  fi
else
  false
fi
