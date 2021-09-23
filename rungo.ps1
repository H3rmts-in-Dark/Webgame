Write-Host "compiling go; starting go `n";

go build -o servermain.exe main.go;
./servermain.exe;