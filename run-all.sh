#!/usr/bin/env bash
set -e

sh personality-test-app-api/run-backend.sh &
sh app/run-frontend.sh &