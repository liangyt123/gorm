@echo off
del gorm.exe
set GOARCH=amd64
set GOOS=linux
go build -o  gorm.exe