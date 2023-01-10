#!/usr/bin/env bash
set -e

# download redis image 
docker pull redis

# build the docker image
docker build -t personality-test-api .

# launch
docker-compose up