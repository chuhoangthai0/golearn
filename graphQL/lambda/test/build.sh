#!/bin/bash

GOOS=linux go build main.go
zip getOne.zip main
rm -rf main