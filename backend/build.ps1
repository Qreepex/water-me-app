# Build script for Water Me backend components

param(
    [string]$Version
)

# Read version from VERSION file if not provided
if ([string]::IsNullOrEmpty($Version)) {
    $Version = Get-Content -Path "VERSION" -Raw | ForEach-Object { $_.Trim() }
}

$Registry = "ghcr.io/qreepex"

Write-Host "Building Water Me backend components v$Version" -ForegroundColor Green

# Build API server
Write-Host "`nBuilding API server..." -ForegroundColor Yellow
docker build --build-arg COMPONENT=api `
  -t "${Registry}/plants-backend-api:${Version}" `
  -t "${Registry}/plants-backend-api:latest" `
  .

# Build notification worker
Write-Host "`nBuilding notification worker..." -ForegroundColor Yellow
docker build --build-arg COMPONENT=notification-worker `
  -t "${Registry}/plants-notification-worker:${Version}" `
  -t "${Registry}/plants-notification-worker:latest" `
  .

Write-Host "`nBuild complete!" -ForegroundColor Green
Write-Host ""
Write-Host "To push to registry, run:" -ForegroundColor Cyan
Write-Host "  docker push ${Registry}/plants-backend-api:${Version}"
Write-Host "  docker push ${Registry}/plants-backend-api:latest"
Write-Host "  docker push ${Registry}/plants-notification-worker:${Version}"
Write-Host "  docker push ${Registry}/plants-notification-worker:latest"
