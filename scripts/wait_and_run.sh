#!/bin/sh
set -e

wait_for_db() {
    echo "Connecting to database ${PG_DB}"
    echo "Connecting to database ${PG_WRITER_HOST}"
    echo "Connecting to database ${PG_USERNAME}"
    echo "Connecting to database ${PG_PASSWORD}"
    echo "Connecting to database ${PG_PORT}"
    until psql postgres://"${PG_USERNAME}":"${PG_PASSWORD}"@"${PG_WRITER_HOST}":"${PG_PORT}"/"${PG_DB}"?sslmode="${PG_SSL_MODE}" -c '\q'; do
      echo >&2 "$(date +%Y%m%dt%H%M%S) Postgres is unavailable - sleeping"
      sleep 3
    done

    echo >&2 "$(date +%Y%m%dt%H%M%S) Postgres is up - starting api.."
}

check_success() {
    # Exit if the last command failed.
    if [ $? -ne 0 ]; then
        echo "Last command failed, exiting.."
        exit 1
    fi
}

set +e

wait_for_db
check_success

set -e
./api
