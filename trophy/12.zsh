#!/usr/bin/env zsh

# title: あれはどこだろな
# desc: cdコマンドを実行した後にlsコマンドを実行した\nあれーどこに有るのかなぁ
# grade: bronze

command grep -e '^ls' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="cd,before"
  if [[ "$ZT_RECORD_HASH[$k]" == "true" ]]; then
    true
  else
    false
  fi
else
  false
fi
