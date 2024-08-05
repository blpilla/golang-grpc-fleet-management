#!/bin/sh
#   Use this script to test if a given TCP host/port are available

set -e

host="$1"
shift
port="$1"
shift

timeout="${WAITFORIT_TIMEOUT:=15}"
cmd="$@"

for i in $(seq $timeout) ; do
    if nc -z "$host" "$port" ; then
        exec $cmd
    fi
    echo "Waiting for $host:$port... ($i/$timeout)"
    sleep 1
done

echo "Operation timed out ($timeout seconds)"
exit 1
