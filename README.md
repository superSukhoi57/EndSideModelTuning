# EndSide Model Tuning

毕业设计项目 - 端侧模型微调系统

## 项目简介

本项目是一个端侧模型微调系统，包含后端服务和前端界面两部分。

## 项目结构

```
EndSideModelTuning/
├── backend/          # 后端服务
│   ├── common/       # 公共模块
│   ├── iterative_control/  # 迭代控制服务
│   └── verify/       # 验证服务
└── frontend/         # 前端界面
```

## 技术栈

- 后端: Go
- 前端: React

## 快速开始

### 后端服务

```bash
cd backend/iterative_control
go run iterative.go
```

### 前端界面

```bash
cd frontend
npm start
```

## 功能特性

- 模型微调
- 迭代控制
- 结果验证



先登录镜像仓库
```bash
docker login your.docker.registry.com
```
然后输入你的用户名和密码
