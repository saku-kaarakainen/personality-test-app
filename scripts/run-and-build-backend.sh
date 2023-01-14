#!/usr/bin/env bash
set -e

# build the docker image
docker build -t personality-test-api .. -f ../api.Dockerfile

# launch in the background
docker-compose up -d personality-test-db personality-test-api
