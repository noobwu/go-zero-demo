@echo off
echo Starting api-front: 9999
pushd %~dp0
go run front.go
popd
pause