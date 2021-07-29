@echo off
echo Starting rpc-sys: 8081
pushd %~dp0
go run ums.go
popd
pause