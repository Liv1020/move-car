#!/usr/bin/env bash
set -e

docker build -t liv1020/move-car-api:pro --no-cache .
docker push liv1020/move-car-api:pro
