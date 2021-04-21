#!/usr/bin/env zsh

# title: 💩💩💩💩💩
# desc: unkoを含むコマンドを5回実行した\n適当な文字列としてunkoを使う癖があるんじゃぁ無いですか？
# grade: bronze

command grep -e 'unko' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="unko,count"
  v=${ZT_RECORD_HASH[$k]:-0}
  if [[ "$v" == "4" ]]; then
    true
  else
    ZT_RECORD_HASH[$k]=$((v+1))
    false
  fi
else
  false
fi
