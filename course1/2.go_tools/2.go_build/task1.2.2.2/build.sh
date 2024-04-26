#!/bin/bash
printf "Compiling started…\n"
go build -o main main.go
printf "Compiling complete.\n"
printf "Trying to launch program\n"
./main
printf "Program exited.\n"