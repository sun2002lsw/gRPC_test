@echo off
cd /d %~dp0

cd src
go build -o ../bin/client.exe

"../bin/client.exe"
