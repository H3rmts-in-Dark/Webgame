@echo off
echo installing node modules; compiling webassembly + packing typescript and wasm; compiling go; starting go

call npm install
call npm run build

call go build -o servermain.exe main.go
start servermain.exe

exit