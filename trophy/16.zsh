#!/usr/bin/env zsh

# title: 💩💩💩💩💩💩💩💩💩💩
# desc: unkoを含むコマンドを10実行した\nsuper_unkoへのコントリビュートをお待ちしております
# grade: bronze

command grep -e 'unko' &>/dev/null
if [[ "$?" == "0" ]]; then
  k="unko,count"
  v=${ZT_RECORD_HASH[$k]:-0}
  if [[ "$v" == "10" ]]; then
    true
  else
    ZT_RECORD_HASH[$k]=$((v+1))
    false
  fi
else
  false
fi
