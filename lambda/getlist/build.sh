#!/bin/bash

GOOS=linux go build getlist.go
zip getlist.zip getlist
rm -rf getlist