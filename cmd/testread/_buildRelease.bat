@echo off
del testread.exe
go build -trimpath -ldflags "-s -w" .