k8s 的 Ingress **可以跨 Namespace 路由**，能随便转发到「其他命名空间」的 Service。

------

## 一、核心原理

Ingress 资源本身**属于某个命名空间**，

但 Ingress 转发后端 Service 时，支持两种写法：

### 1. 同命名空间（简写，只写服务名）

```
backend:
  service:
    name: demo-svc
    port:
      number: 80
```

默认找 **当前 Ingress 所在 ns** 的服务。

### 2. 跨命名空间（完整写法：服务名 + 命名空间）

```
backend:
  service:
    name: demo-svc
    namespace: test-ns   # 关键：指定其他命名空间
    port:
      number: 80
```

✅ 直接跨 ns 路由，**生产非常常用**。