#!/usr/bin/env zsh

# title: はじめてのrm
# desc: はじめてrmコマンドを実行した\nバイバイファイル
# grade: bronze

command grep -e '^rm ' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="rm,count"
  val=${ZT_RECORD_HASH[$k]:-0}
  if [[ "$val" == "0" ]]; then
    true
  else
    false
  fi
else
  false
fi
