@echo off
echo Starting rpc-sys: 8080
pushd %~dp0
go run sys.go
popd
pause