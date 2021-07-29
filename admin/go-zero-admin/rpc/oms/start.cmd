@echo off
echo Starting rpc-oms: 8083
pushd %~dp0
go run oms.go
popd
pause