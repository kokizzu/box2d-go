#!/bin/sh

set -e -u -o pipefail

cd transpile
go run . "$@"
