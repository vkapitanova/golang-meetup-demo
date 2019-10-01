#!/bin/bash

export LC_ALL=en_US.utf-8
export LANG=en_US.utf-8

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/simple-docker .
