@echo off
echo Starting rpc-sms: 8084
pushd %~dp0
go run sms.go
popd
pause