#!/bin/bash

# Usage details:
# ./setup-db.sh <DB_HOST> <DB_PORT> <DB_ROOT_PASSWORD> <DB_USER> <DB_PASSWORD> <DB_NAME>
# Meant to be run once per postgres server

if [ $# -lt 6 ]; then
  echo 1>&2 "not enough arugments!\nusage ./setup-db.sh <db_host> <db_port> <db_root_password> <db_user> <db_password> <db_name>"
  exit 2
elif [ $# -gt 6 ]; then
  echo 1>&2 "too many arugments!\nusage ./setup-db.sh <db_host> <db_port> <db_root_password> <db_user> <db_password> <db_name>"
  exit 2
fi

DB_HOST="$1"
DB_PORT="$2"
DB_USER="$4"
DB_PASSWORD="$5"
DB_NAME="$6"

export PGPASSWORD=$3

psql -h $DB_HOST -p $DB_PORT -U postgres -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';"
psql -h $DB_HOST -p $DB_PORT -U postgres -c "CREATE DATABASE $DB_NAME OWNER $DB_USER;"
psql -h $DB_HOST -p $DB_PORT -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
