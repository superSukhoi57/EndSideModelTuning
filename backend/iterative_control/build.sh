#!/bin/bash

set -e

VERSION_FILE="$(dirname "$0")/version.txt"
MAX_VERSION_RECORDS=88
CONFIG_FILE="$(dirname "$0")/../../build-config.json"

echo "=== Building iterative control service ==="

# 读取当前版本号
if [ -f "$VERSION_FILE" ] && [ -s "$VERSION_FILE" ]; then
    CURRENT_VERSION=$(head -n 1 "$VERSION_FILE")
else
    CURRENT_VERSION="0.0.0"
    echo "未找到版本记录，使用初始版本: $CURRENT_VERSION"
fi

echo "当前版本: $CURRENT_VERSION"

# 解析版本号
IFS='.' read -r X1 X2 X3 <<< "$CURRENT_VERSION"

# 提示用户选择递增位置
echo ""
echo "请选择版本号递增位置:vX1.X2.X3"
echo "  1 - x1 主版本（模块重构或大变更）"
echo "  2 - x2 次要版本（功能迭代）"
echo "  3 - x3 补丁版本（小改动、bug修复）"
read -r -p "请输入 [1/2/3]: " BUMP_CHOICE

case "$BUMP_CHOICE" in
    1)
        X1=$((X1 + 1))
        X2=0
        X3=0
        ;;
    2)
        X2=$((X2 + 1))
        X3=0
        ;;
    3)
        X3=$((X3 + 1))
        ;;
    *)
        echo "错误: 无效输入，请输入 1、2 或 3"
        exit 1
        ;;
esac

NEW_VERSION="${X1}.${X2}.${X3}"
echo "新版本: $NEW_VERSION"
echo ""

# 构建Docker镜像
echo "Building Docker image: iterative:$NEW_VERSION ..."
docker build -t "iterative:$NEW_VERSION" -f Dockerfile ..

echo "=== Docker image built successfully ==="

# 推送镜像到远程仓库（从配置文件读取）
REGISTRY_BASE=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['remote_registry'])" 2>/dev/null || python -c "import json; print(json.load(open('$CONFIG_FILE'))['remote_registry'])")
REMOTE_REGISTRY="${REGISTRY_BASE}/iterative"
REMOTE_IMAGE="${REMOTE_REGISTRY}:${NEW_VERSION}"

echo "推送镜像到远程仓库: $REMOTE_IMAGE ..."
docker tag "iterative:${NEW_VERSION}" "$REMOTE_IMAGE"
docker push "$REMOTE_IMAGE"
echo "镜像推送成功"

# 删除远程仓库标签引用，保留本地标签
echo "清理远程仓库标签引用..."
docker rmi "$REMOTE_IMAGE"
echo "已删除标签引用: $REMOTE_IMAGE"
echo "保留本地标签: iterative:${NEW_VERSION}"

# 构建成功后，更新版本记录
echo "更新版本记录..."

# 创建临时文件，先写入新版本号
TEMP_FILE=$(mktemp)
echo "$NEW_VERSION" > "$TEMP_FILE"

# 如果原文件存在，追加旧版本（最多保留 MAX_VERSION_RECORDS - 1 条）
if [ -f "$VERSION_FILE" ] && [ -s "$VERSION_FILE" ]; then
    head -n $((MAX_VERSION_RECORDS - 1)) "$VERSION_FILE" >> "$TEMP_FILE"
fi

# 替换原文件
mv "$TEMP_FILE" "$VERSION_FILE"

echo "版本记录已更新，当前保留 $(wc -l < "$VERSION_FILE") 条记录"
echo "最新版本: $NEW_VERSION"
