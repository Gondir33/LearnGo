#!/bin/bash
go build -o myprogram $1
echo "Debug started..."
dlv exec ./myprogram
echo "Debug ended."