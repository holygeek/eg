#!/bin/sh
GOPATH=`pwd`
export GOPATH

eg -w -t template.go foo
