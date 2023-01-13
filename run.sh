#!/usr/bin/env bash
set -e

# build the docker image
docker build -t personality-test-api .
docker build -t personality-test-app .

# launch
docker-compose up