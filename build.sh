#!/usr/bin/env bash 

# install packages and dependencies
go get github.com/gin-gonic/gin

go get github.com/tpkeeper/gin-dump

go get gopkg.in/go-playground/validator.v9

# build command
go build -o bin/application server.go