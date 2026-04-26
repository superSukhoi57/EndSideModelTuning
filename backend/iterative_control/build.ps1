chcp 65001 > $null

$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$VERSION_FILE = Join-Path $ScriptDir "version.txt"
$MAX_VERSION_RECORDS = 88
$CONFIG_FILE = Join-Path $ScriptDir "../../build-config.json" | Resolve-Path

Write-Host "=== Building iterative control service ==="

if ((Test-Path $VERSION_FILE) -and (Get-Item $VERSION_FILE).Length -gt 0) {
    $CURRENT_VERSION = Get-Content $VERSION_FILE -TotalCount 1
} else {
    $CURRENT_VERSION = "0.0.0"
    Write-Host "未找到版本记录，使用初始版本: $CURRENT_VERSION"
}

Write-Host "当前版本: $CURRENT_VERSION"

$X1, $X2, $X3 = $CURRENT_VERSION -split '\.' | ForEach-Object { [int]$_ }

Write-Host ""
Write-Host "请选择版本号递增位置:vX1.X2.X3"
Write-Host "  1 - x1 主版本（模块重构或大变更）"
Write-Host "  2 - x2 次要版本（功能迭代）"
Write-Host "  3 - x3 补丁版本（小改动、bug修复）"
$BUMP_CHOICE = Read-Host "请输入 [1/2/3]"

switch ($BUMP_CHOICE) {
    "1" {
        $X1 = $X1 + 1
        $X2 = 0
        $X3 = 0
    }
    "2" {
        $X2 = $X2 + 1
        $X3 = 0
    }
    "3" {
        $X3 = $X3 + 1
    }
    default {
        Write-Host "错误: 无效输入，请输入 1、2 或 3"
        exit 1
    }
}

$NEW_VERSION = "${X1}.${X2}.${X3}"
Write-Host "新版本: $NEW_VERSION"
Write-Host ""

Write-Host "Building Docker image: iterative:$NEW_VERSION ..."
docker build -t "iterative:$NEW_VERSION" .

Write-Host "=== Docker image built successfully ==="

$ConfigContent = Get-Content $CONFIG_FILE -Raw | ConvertFrom-Json
$REGISTRY_BASE = $ConfigContent.remote_registry
$REMOTE_REGISTRY = "${REGISTRY_BASE}/iterative"
$REMOTE_IMAGE = "${REMOTE_REGISTRY}:${NEW_VERSION}"

Write-Host "推送镜像到远程仓库: $REMOTE_IMAGE ..."
docker tag "iterative:${NEW_VERSION}" "$REMOTE_IMAGE"
docker push "$REMOTE_IMAGE"
Write-Host "镜像推送成功"

Write-Host "清理远程仓库标签引用..."
docker rmi "$REMOTE_IMAGE"
Write-Host "已删除标签引用: $REMOTE_IMAGE"
Write-Host "保留本地标签: iterative:${NEW_VERSION}"

Write-Host "更新版本记录..."

$TEMP_FILE = [System.IO.Path]::GetTempFileName()
$NEW_VERSION | Out-File -FilePath $TEMP_FILE -Encoding utf8

if ((Test-Path $VERSION_FILE) -and (Get-Item $VERSION_FILE).Length -gt 0) {
    Get-Content $VERSION_FILE | Select-Object -First ($MAX_VERSION_RECORDS - 1) | Add-Content -Path $TEMP_FILE
}

Move-Item -Path $TEMP_FILE -Destination $VERSION_FILE -Force

$LINE_COUNT = (Get-Content $VERSION_FILE | Measure-Object -Line).Lines
Write-Host "版本记录已更新，当前保留 $LINE_COUNT 条记录"
Write-Host "最新版本: $NEW_VERSION"
