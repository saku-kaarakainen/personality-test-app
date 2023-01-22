#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

# TODO: Convert this into Makefile?

# build the docker image
docker build -t personality-test-api . 

# launch
docker-compose up