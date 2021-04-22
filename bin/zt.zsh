#!/usr/bin/env zsh

ZT_DIR=$(cd $(dirname $0); cd ../; pwd)
ZT_RECORD=${ZT_RECORD_JSON:-$ZT_DIR/.record/$USER}

[[ -e $ZT_DIR/.record ]] || mkdir -p $ZT_DIR/.record
touch $ZT_RECORD
typeset -g -A ZT_RECORD_HASH=()

function _zt_load_record() {
  [[ ! -e "$ZT_RECORD" ]] && ZT_RECORD_HASH=() && return 0
  command sed -E 's/([^:]*): *(.*)/\1 \2/' $ZT_RECORD | while read key value; do
    ZT_RECORD_HASH[$key]=$value
  done
}

_zt_load_record

function _zt-main() {
  zmodload zsh/zutil
  local -A opthash
  zparseopts -D -A opthash -- h -help || return 1

  autoload -Uz add-zsh-hook
  add-zsh-hook -Uz preexec zt-preexec
  add-zsh-hook -Uz zshexit zt-zshexit
}

function zt-preexec() {
  local cmd="${1}"

  _zt_torophy_checker_list | while read T; do
    local rn="$(basename $T)"
    rn="${rn%.zsh}"
    if [[ "$ZT_RECORD_HASH[$rn]" == "false" ]] || [[ "$ZT_RECORD_HASH[$rn]" == "" ]]; then
      ZT_RECORD_HASH[$rn]="false"
      if echo "'$cmd'" | source $T &>/dev/null; then
        ZT_RECORD_HASH[$rn]="true"
        _zt_print_trophy $T
      fi
    fi
  done
}

function zt-zshexit() {
  zt-save
}

function zt-save() {
  [[ "$ZT_FLUSHING" == "1" ]] && return
  ZT_FLUSHING=1
  for k in ${(k)ZT_RECORD_HASH}; do
    echo "$k: $ZT_RECORD_HASH[$k]"
  done > $ZT_RECORD
  ZT_FLUSHING=0
}

function _zt_is_cleared_trophy {

  echo is_cleared: "["$ZT_RECORD_HASH[$1]"]" >&2

  if [[ "$ZT_RECORD_HASH[$1]" == "" ]]; then
    ZT_RECORD_HASH[${1}]="false"
    echo 0
  elif [[ "$ZT_RECORD_HASH[$1]" == "false" ]]; then
    echo 0
  else 
    echo 1
  fi
  echo is_cleared-after: "["$ZT_RECORD_HASH[$1]"]" >&2
}

function _zt_print_trophy() {
  local target="${1}"
  local title=$(command grep -m1 '^# title: ' $target | command sed -E 's/[^:]*: *(.*)/\1/')
  local  desc=$(command grep -m1 '^# desc: '  $target | command sed -E 's/[^:]*: *(.*)/\1/')
  local grade=$(command grep -m1 '^# grade: ' $target | command sed -E 's/[^:]*: *(.*)/\1/')
  typeset -A emoji=("bronze" "\U1F949" "silver" "\U1F948" "gold" "\U1F947" "special" "\U1F3C6")

  local width=$({ echo -e $title; echo -e $desc } | while read L;do echo ${#L}; done | command sort -nr | command head -n1)

  (
    echo ""
    echo "${emoji[special]} zsh-trophy ${emoji[special]}"
    printf "%*s\n" ${COLUMNS:-$(tput cols)} '' | tr ' ' '.'
    echo -e "..\t\033[48;2;66;66;66m  ${emoji[$grade]} $title  \033[0m"
    echo -e "${desc}" | while read L;do
      local length=${#L}
      echo -en "..\t\t\033[48;2;66;66;66m"
      printf "%*s" $(((width-length)/2)) " "
      echo -en "$L"
      printf "%*s" $(((width-length)/2)) " "
      echo -e '\033[0m'
    done
    printf "%*s\n" ${COLUMNS:-$(tput cols)} '' | tr ' ' '.'
    echo ""
  ) >&2
}

function _zt_torophy_checker_list() {
  command ls $ZT_DIR/trophy/*.zsh
}

_zt-main
