#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

scp main LGWallet:/data/api
