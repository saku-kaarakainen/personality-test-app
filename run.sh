#!/usr/bin/env bash
set -e


# TODO: Convert this into Makefile?

# build the docker image
docker build -t personality-test-api . -f api.Dockerfile
docker build -t personality-test-app . -f app.Dockerfile

# launch
docker-compose up