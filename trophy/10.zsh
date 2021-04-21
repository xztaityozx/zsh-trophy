#!/usr/bin/env zsh

# title: 自由自在
# desc: cdコマンドを3回実行した
# grade: bronze

command grep -e '^cd' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="cd,count"
  val=${ZT_RECORD_HASH[$k]:-0}
  if [[ "$val" == "2" ]]; then
    true
  else
    ZT_RECORD_HASH[$k]=$(($val+1))
    false
  fi
else
  false
fi
