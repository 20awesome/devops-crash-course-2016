#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
sudo docker build -t vitaliihurin/test-service .