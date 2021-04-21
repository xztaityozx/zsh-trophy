#!/usr/bin/env zsh

# title: 3回目のls
# desc: lsコマンドを3回叩いた\ncdしたあとってlsしちゃうよね。なんでだろうね
# grade: bronze

command grep -e '^ls' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="ls,count"
  if [[ "$ZT_RECORD_HASH[$k]" == "2" ]]; then
    true
  else
    false
  fi
else
  false
fi

