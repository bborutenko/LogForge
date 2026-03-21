#!/bin/bash

cd "$(dirname "$0")/../../../apps/api/docker/postgres" || exit

docker build -t logforge-postgres .

docker rm -f logforge-db 2>/dev/null || true

docker run -d \
  --name logforge-db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=logforge \
  -p 5432:5432 \
  logforge-postgres
