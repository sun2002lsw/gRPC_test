@echo off
cd /d %~dp0

protoc sayHello.proto --go_out=plugins=grpc:./Server/pkg/
protoc sayHello.proto --go_out=plugins=grpc:./Client_go/pkg/