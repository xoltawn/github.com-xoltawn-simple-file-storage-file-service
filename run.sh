#! /bin/bash

cp .env.example .env
cat .env
docker network create --driver bridge simple-file-storage || true
docker compose up