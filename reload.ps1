$params = @{
action = "Reload Site"
code = "supersaveandsecureAPIcode"
}
Invoke-WebRequest -Uri "http://localhost:18265/api" -Body ($params | ConvertTo-Json) -Method POST -UseBasicParsing