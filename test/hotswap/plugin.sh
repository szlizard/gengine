#!/bin/bash

# go install github.com/edwingeng/hotswap/cli/hotswap@latest
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 hotswap build . .
# mv -f ./hotswap2.so ./hotswap_M_m.so
