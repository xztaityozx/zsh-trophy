#!/usr/bin/env zsh

# title: 50回目のls
# desc: lsコマンドを50回実行した\nもう手が勝手に動く
# grade: silver

command grep -e '^ls' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="ls,count"
  if [[ "$ZT_RECORD_HASH[$k]" == "49" ]]; then
    true
  else
    false
  fi
else
  false
fi
