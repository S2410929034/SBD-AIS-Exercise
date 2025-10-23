#!/bin/sh

# todo
# docker build
docker build . -t drink-api:latest

docker network create service-net

# docker run db
docker run -d \
    --name drink-db \
    --env-file debug.env \
    -v drinkdbdata:/var/lib/postgresql/18/docker \
    --network service-net \
    -p 5432:5432 \
    postgres:18

sleep 5

# docker run orderservice
docker run -d \
    --name drink-webserver \
    --env-file debug.env \
    --network service-net \
    -p 3000:3000 \
    drink-api:latest