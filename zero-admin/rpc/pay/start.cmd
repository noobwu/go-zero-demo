@echo off
echo Starting rpc-pay: 8085
pushd %~dp0
go run pay.go
popd
pause