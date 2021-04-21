#!/usr/bin/env zsh

# title: パパパパイプライン
# desc: パイプが4個含まれるコマンドを実行した
# grade: bronze

[[ "$(command grep -o . | command grep -c '|')" == "4" ]]
