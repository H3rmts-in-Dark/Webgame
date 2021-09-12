# Webgame
 Webgame 

Cmd to test API
> cmd
> curl http://192.168.187.10:18265/api -d "{\"action\":\"Reload Site\",\"code\":\"supersaveandsecureAPIcode\"}"


Powershell:
> $params = @{
> action = "Reload Site"
> code = "supersaveandsecureAPIcode"
> }
> Invoke-WebRequest -Uri "http://server-debian.local:18265/api" -Body ($params | ConvertTo-Json) -Method POST