#!/bin/bash

set -e

go get github.com/kevinburke/go-bindata

# remove generated files first
rm bindata.go 2>/dev/null || true

# then re-create them
find . -name '*.css' -or -name '*.png' -or -name '*.js' -or -name '*.svg' -or -name '*.html' | \
    xargs go-bindata -pkg=standalone
