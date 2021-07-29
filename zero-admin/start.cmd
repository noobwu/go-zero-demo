@echo off
echo Starting RPC: 
pushd %~dp0\rpc
@start "rpc-sys: 8080" /min "%~dp0\rpc\sys\start.cmd"
@start "rpc-ums: 8081" /min "%~dp0\rpc\ums\start.cmd"
@start "rpc-pms: 8082" /min "%~dp0\rpc\pms\start.cmd"
@start "rpc-oms: 8083" /min "%~dp0\rpc\oms\start.cmd"
@start "rpc-sms: 8084" /min "%~dp0\rpc\sms\start.cmd"
@start "rpc-pay: 8085" /min "%~dp0\rpc\pay\start.cmd"
popd

echo Starting api: 
@start "api-admin: 8888" /min "%~dp0\api\start.cmd"
@start "api-front: 9999" /min "%~dp0\front-api\start.cmd"