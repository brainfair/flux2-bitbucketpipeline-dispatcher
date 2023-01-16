$filename = "test.json"
$url="http://localhost:8000/webhook"

$content = gc $filename

[hashtable]$headers=@{}

$headers.Add('Content-Type', 'application/json')

$statusCode = Invoke-WebRequest -Uri $url -Method POST -Body $content -Headers $headers

Write-Host $statusCode