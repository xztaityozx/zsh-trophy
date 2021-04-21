#!/usr/bin/env zsh

# title: 1000回目のls
# desc: lsコマンドを1000回実行した\nキミは実質lsコマンドだな
# grade: special

command grep -e '^ls' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="ls,count"
  if [[ "$ZT_RECORD_HASH[$k]" == "999" ]]; then
    true
  else
    ZT_RECORD_HASH[$k]=$((${ZT_RECORD_HASH[$k]:-0}+1))
    false
  fi
else
  false
fi
