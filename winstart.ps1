Write-Host "installing node modules; compiling webassembly + packing typescript and wasm; `n"

npm install;
npm run build;

Start-Process PowerShell ./rungo.ps1