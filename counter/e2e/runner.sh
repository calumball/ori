#!/bin/bash
# e2e test for the counter service

echo "Building counter service"
docker build . -t proto -f proto.Dockerfile
docker build -t ori-app ./counter

echo "Starting counter server"
docker run -d --rm -it -p 8888:8888 --name test-server ori-app

echo "Starting test suite"
go test counter/e2e/e2e_test.go --addr localhost:8888

echo "Cleaning up test server"
docker kill test-server