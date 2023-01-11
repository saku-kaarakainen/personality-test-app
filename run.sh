#!/usr/bin/env bash
set -e

# build the docker image
docker build -t personality-test-api .

# launch
docker-compose up