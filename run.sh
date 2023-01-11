#!/usr/bin/env bash
set -e

# bootstrap API
docker build -t personality-test-api .

# launch API
docker run -it --rm -p 8080:8080 personality-test-api