@echo off
del .\gorm-tool
set GOARCH=amd64
set GOOS=linux
go build -o  gorm