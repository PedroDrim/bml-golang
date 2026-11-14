#!/bin/bash
go build start.go
./start config.json > /dev/null
rm start
cat benchmark.json