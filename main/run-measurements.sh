#!/usr/bin/bash

go run *.go -measure Bnb &\
go run *.go -measure NetworkFlow &\
go run *.go -measure Crown