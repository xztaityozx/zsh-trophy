#!/usr/bin/env zsh

# title: 200回目のls
# desc: lsコマンドを200回実行した\nlsキーがキーボードに欲しいぐらいだよ
# grade: gold

command grep -e '^ls' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="ls,count"
  if [[ "$ZT_RECORD_HASH[$k]" == "199" ]]; then
    true
  else
    false
  fi
else
  false
fi
