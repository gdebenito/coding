#!/bin/sh
go run main.go < in > res  && diff out res
