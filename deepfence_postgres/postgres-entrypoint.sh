#!/bin/bash

echoerr() { echo "$@" 1>&2; }

export POSTGRES_MULTIPLE_DATABASES="$DEEPFENCE_POSTGRES_USER_DB_NAME"
export POSTGRES_USER=$DEEPFENCE_POSTGRES_USER_DB_USER
export POSTGRES_PASSWORD=$DEEPFENCE_POSTGRES_USER_DB_PASSWORD
export PGPASSWORD=$DEEPFENCE_POSTGRES_USER_DB_PASSWORD

/bin/bash /usr/local/bin/create-pg-dirs.sh && /bin/bash /usr/local/bin/new-docker-entrypoint.sh postgres -c shared_buffers=256MB -c max_connections=1000

#sleep 30
#gosu postgres pg_ctl stop
#sleep 10
#cp /usr/local/postgresql.conf $PGDATA/postgresql.conf
#chown postgres:postgres $PGDATA/postgresql.conf
#/bin/bash /usr/local/bin/new-entry-point.sh postgres
