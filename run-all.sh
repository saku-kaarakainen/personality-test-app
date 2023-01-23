#!/usr/bin/env bash
set -e

sh api/run-backend.sh &
sh app/run-frontend.sh &