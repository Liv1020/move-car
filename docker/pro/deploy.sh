#!/usr/bin/env bash
set -e

cd /var/web
docker-compose pull move-car-api
docker-compose up --force-recreate --no-deps -d move-car-api