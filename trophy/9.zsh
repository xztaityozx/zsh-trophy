#!/usr/bin/env zsh

# title: はじめてのcd
# desc: cdコマンドを初めて実行した\nなにはともあれcdコマンド
# grade: bronze

command grep -e '^cd' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="cd,count"
  val=${ZT_RECORD_HASH[$k]:-0}
  if [[ "$val" == "0" ]]; then
    true
  else
    false
  fi
else
  false
fi
