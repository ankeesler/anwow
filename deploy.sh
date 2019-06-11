#!/usr/bin/env bash

set -exou pipefail

dir="$(mktemp -d)"
GOOS=linux go build -o "$dir/anwow" main.go
cf push anwow -b binary_buildpack -c './anwow' -p "$dir"
rm -rf "$dir"
