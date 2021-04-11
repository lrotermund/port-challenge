#!/bin/sh
echo "go lint :: started"
golangci-lint run ./... --out-format checkstyle > report.xml -v
echo "go lint :: finished"