#!/bin/bash

CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -trimpath -buildmode=plugin -o=plugin_M_m.so plugin_superman.go