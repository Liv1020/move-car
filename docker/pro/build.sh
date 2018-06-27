#!/usr/bin/env bash
set -e

docker build -t liv1020/go-move-car:pro --no-cache .
docker push liv1020/go-move-car:pro
