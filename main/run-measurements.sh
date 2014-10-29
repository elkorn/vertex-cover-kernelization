#!/usr/bin/bash

go run *.go -run=Bnb &\
go run *.go -run=Naive &\
go run *.go -run=NetworkFlow &\
go run *.go -run=Crown