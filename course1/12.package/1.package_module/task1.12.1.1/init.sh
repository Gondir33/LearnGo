#!/bin/bash
if [ -z "$1" ]; then
    echo "Module name argument is missing"
    exit 1
fi
    go mod init $1