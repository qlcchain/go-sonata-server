#!/usr/bin/env bash

goimports -w -local github.com/qlcchain/go-sonata-server $(find . -type f -name '*.go' -not -path "*/mocks/*" -not -path "*/pb/*")
