#!/usr/bin/env zsh

ZT_DIR=$(cd $(dirname $0)/../; pwd)

function _zsh-trophy-preexec() {
  [[ ! -e "$ZT_DIR/bin/zsh-trophy" ]] && return 0

  $ZT_DIR/bin/zsh-trophy --ztd=$ZT_DIR --cmd="${1}" --width="${COLUMNS:-$(tput cols)}"
}

autoload -Uz add-zsh-hook

add-zsh-hook -Uz preexec _zsh-trophy-preexec
