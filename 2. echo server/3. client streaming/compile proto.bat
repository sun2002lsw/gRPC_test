@echo off
cd /d %~dp0

protoc echo.proto --go_out=plugins=grpc:./Server/pkg/
protoc echo.proto --go_out=plugins=grpc:./Client/pkg/