#!/bin/sh

protoc --go_out=/go/src -I=./gcommand ./gcommand/gcommand.proto

protoc --go-grpc_out=/go/src -I=./gcommand ./gcommand/gcommand.proto
