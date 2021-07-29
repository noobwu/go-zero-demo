@echo off
echo Starting api-admin: 8888
pushd %~dp0
go run admin.go
popd
pause