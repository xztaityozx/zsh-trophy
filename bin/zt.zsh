#!/usr/bin/env zsh

ZT_DIR=$(cd $(dirname $0); cd ../; pwd)
ZT_RECORD_JSON=${ZT_RECORD_JSON:-$ZT_DIR/.record/$USER.json}

function main() {
  zmodload zsh/zutil
  local -A opthash
  zparseopts -D -A opthash -- h -help || return 1

  if ! type jq &> /dev/null; then
    echo [zsh-tropy] jq is required >&2
    return 0
  fi

  autoload -Uz add-zsh-hook
  add-zsh-hook -Uz preexec zt-preexec

  [[ -e $ZT_DIR/.record ]] || mkdir -p $ZT_DIR/.record
  touch $ZT_RECORD_JSON
}

function zt-preexec() {
  local cmd="${1}"

  _zt_torophy_checker_list | while read T; do
    _zt_is_cleared_trophy $T
    #echo "$cmd" | zsh $T &>/dev/null
    #if [[ "$?" == "0" ]] then
      #_zt_print_trophy $T
    #fi
  done
}

function _zt_is_cleared_trophy() {
  local t="$(basename $1)"
  local q=".\\\"${t%.zsh}\\\" | .cleared"

  if [[ "$(jq -r $q)" == "true" ]] then
    echo 1
  fi

  echo 0
}

function _zt_print_trophy() {
  local target="${1}"
  local title=$(command grep -m1 '^# title: ' $target | command cut -d: -f 2-| command sed 's/^ *//g')
  local  desc=$(command grep -m1 '^# desc: '  $target | command cut -d: -f 2-| command sed 's/^ *//g')
  local grade=$(command grep -m1 '^# grade: ' $target | command cut -d: -f 2-| command sed 's/^ *//g')
  typeset -A emoji=("bronze" "\U1F949" "silver" "\U1F948" "gold" "\U1F947" "special" "\U1F3C6")

  local width=$({ echo -e $title; echo -e $desc } | command awk '$1=length' | command sort -nr | command head -n1)

  echo ""
  echo "${emoji[special]} zsh-trophy ${emoji[special]}"
  printf "%*s\n" ${COLUMNS:-$(tput cols)} '' | tr ' ' '.'
  echo -e "..\t\033[48;2;66;66;66m  ${emoji[$grade]} $title  \033[0m"
  echo -e "${desc}" | while read L;do
    local length=$(echo -en $L | awk '$1=length')
    echo -en "..\t\t\033[48;2;66;66;66m"
    printf "%*s" $(((width-length)/2)) " "
    echo -en "$L"
    printf "%*s" $(((width-length)/2)) " "
    echo -e '\033[0m'
  done
  printf "%*s\n" ${COLUMNS:-$(tput cols)} '' | tr ' ' '.'
  echo ""
}

function _zt_torophy_checker_list() {
  command ls $ZT_DIR/trophy/*.zsh
}

main
