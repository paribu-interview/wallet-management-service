#!/usr/bin/env bash
set -e

host="$1"
port="$2"
shift 2

while ! nc -z "$host" "$port"; do
  >&2 echo "Waiting for $host:$port to become available..."
  sleep 1
done

>&2 echo "$host:$port is available, proceeding..."
exec "$@"
