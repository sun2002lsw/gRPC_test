@echo off
cd /d %~dp0

cd src
go build -o ../bin/server.exe

"../bin/server.exe"
