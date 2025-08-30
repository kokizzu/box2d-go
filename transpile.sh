#!/usr/bin/env bash

set -e -u -o pipefail

cd transpile
go run . "$@"
