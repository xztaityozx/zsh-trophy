#!/usr/bin/env zsh

# title: 10回目のls
# desc: lsコマンドを10回叩いた\n意味もなくlsコマンド実行することってあるよね
# grade: bronze

command grep -e '^ls' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="ls,count"
  if [[ "$ZT_RECORD_HASH[$k]" == "9" ]]; then
    true
  else
    false
  fi
else
  false
fi
