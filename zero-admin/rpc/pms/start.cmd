@echo off
echo Starting rpc-pms: 8082
pushd %~dp0
go run pms.go
popd
pause