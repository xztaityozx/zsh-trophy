#!/usr/bin/env zsh

# title: どこいくの？
# desc: cdコマンドを連続で実行した\ncdしたあとに「ああ！ここじゃないよ！」ってなること有るよね
# grade: bronze

command grep -e '^cd' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="cd,before"
  if [[ "$ZT_RECORD_HASH[$k]" == "true" ]]; then
    true
  else
    ZT_RECORD_HASH[$k]="true"
    false
  fi
else
  ZT_RECORD_HASH[$k]="false"
  false
fi
