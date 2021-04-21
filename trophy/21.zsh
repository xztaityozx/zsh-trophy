#!/usr/bin/env zsh

# title: rmマスター
# desc: rmコマンドを50回実行した\n過去の嫌な思い出もrmしたい
# grade: gold

command grep -e '^rm ' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="rm,count"
  val=${ZT_RECORD_HASH[$k]:-0}
  if [[ "$val" == "49" ]]; then
    true
  else
    ZT_RECORD_HASH[$k]=$(($val+1))
    false
  fi
else
  false
fi
