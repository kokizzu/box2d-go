#!/usr/bin/env bash

set -e -u -o pipefail

rm -rf box2d-c
git clone --rev v3.1.1 https://github.com/erincatto/box2d.git box2d-c

git -C box2d-c am < box2d-3.1.1.patch
