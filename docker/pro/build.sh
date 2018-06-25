#!/usr/bin/env bash
set -e

docker build -t liv1020/go-move-car:pro .
docker push liv1020/go-move-car:pro
