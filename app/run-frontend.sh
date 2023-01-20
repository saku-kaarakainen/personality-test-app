#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

# TODO: Convert this into Makefile?

# build the docker image
docker build -t personality-test-app .

# launch
docker-compose up