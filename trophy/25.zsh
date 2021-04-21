#!/usr/bin/env zsh

# title: 立派なワンライナー
# desc: パイプが10個含まれるコマンドを実行した
# grade: silver

[[ "$(command grep -o . | command grep -c '|')" == "10" ]]
