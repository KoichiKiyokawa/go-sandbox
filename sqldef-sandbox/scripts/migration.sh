#!/bin/bash

set -euo pipefail
cd "$(dirname "$0")"

HOST="db"
DB_NAME="postgres"
DB_USER="postgres"
DB_PASSWORD="postgres"

MERGED_SCHMEA=$(cat ../ddl/*)

# INSERT文を削除(INSERT<任意の文字列の連続>;の形式を削除)
# ダブルクォーテーションで囲わないと、変数に入れたときに改行がなくなってしまうことに注意 https://ex1.m-yabe.com/archives/3339
INSERT_REMOVED_SCHEMA=$(echo "$MERGED_SCHMEA" | perl -0pe 's/INSERT.+?;//gs')

COMMAND='echo "$INSERT_REMOVED_SCHEMA" | psqldef -h $HOST -U $DB_USER $DB_NAME -W $DB_PASSWORD'

eval $COMMAND --dry-run
while true; do
  printf "実行しますか？(y/N)\n>"
  read answer
  case $answer in
    [Yy] ) eval $COMMAND; break;; # Y or y
    [Nn""] ) break;; # N or n or Enter
    * ) : # other
  esac
done
