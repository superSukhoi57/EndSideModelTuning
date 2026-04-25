并非所有使用 Helm 部署的应用都是有状态的。应用是否有状态，以及重启后数据是否会丢失，取决于应用本身的设计和其 Helm Chart 的具体配置。

简单来说，Helm 只是一个 Kubernetes 的包管理工具，它负责按照你给定的配置来部署应用。应用是否有状态，关键在于它是否依赖持久化存储来保存数据。

### 📦 无状态应用 vs. 有状态应用

在 Kubernetes 中，应用通常被分为两类：

- **无状态应用 (Stateless Applications)**
  - **特点**：应用实例不保存任何持久化的数据或状态。每个请求都可以由任何一个实例处理，实例之间是完全相同且可以互换的。
  - **典型代表**：Web 服务器（如 Nginx）、API 网关等。
  - **重启影响**：重启后数据**不会丢失**，因为它们本来就不在本地存储数据。
- **有状态应用 (Stateful Applications)**
  - **特点**：应用实例需要保存持久化的数据或状态。每个实例都有自己独特的身份和存储。
  - **典型代表**：数据库（如 MySQL, PostgreSQL）、消息队列（如 Kafka）等。
  - **重启影响**：如果配置不当，重启后数据**会丢失**。

### 🗃️ 数据持久化：关键在配置

对于有状态应用，数据是否会丢失，核心在于 Helm Chart 中是否配置了**持久化存储**。

1. **数据会丢失的情况**
   - 如果 Helm Chart 默认使用临时存储（如 `emptyDir`），或者你在安装时没有启用持久化选项，那么所有数据都存储在 Pod 的生命周期内。一旦 Pod 被删除或重启，数据就会永久丢失。
   - 例如，部署一个 MySQL 数据库时，如果没有配置持久化，执行 `helm install my-mysql ... --set primary.persistence.enabled=false`，那么重启后所有数据库数据都会消失。
2. **数据不会丢失的情况**
   - 绝大多数生产级的 Helm Chart 都支持持久化存储。你只需要在安装或升级时，通过 `values.yaml` 文件或 `--set` 参数来启用它。
   - 这通常会涉及到配置 `persistence.enabled=true`、指定 `storageClass` 和存储 `size` 等。
   - 配置后，Helm 会为应用创建**持久卷声明 (PVC)**，Kubernetes 会将其绑定到一个**持久卷 (PV)** 上。这样，即使 Pod 重启，数据也安全地保存在 PV 中，新的 Pod 会重新挂载这个 PV，从而恢复数据。

### 📊 总结对比

| 特性             | 无状态应用 (如 Nginx) | 有状态应用 (如 MySQL)        |
| ---------------- | --------------------- | ---------------------------- |
| **数据持久化**   | 不需要                | **必需**，否则重启会丢数据   |
| **重启后数据**   | 不受影响              | 配置持久化则保留，否则丢失   |
| **Helm配置关键** | 通常无需特殊配置      | 需启用 `persistence.enabled` |

因此，在部署任何应用前，尤其是数据库等有状态服务，务必检查其 Helm Chart 的文档和默认配置，确保正确设置了持久化存储，以保障数据安全。

### 🛡️ 关键配置：`resourcePolicy: "keep"`

Harbor 的 Helm Chart 中有一个非常重要的配置项，专门用来防止数据被误删：

在 `values.yaml` 文件中：

```yaml
persistence:
  enabled: true
  resourcePolicy: "keep"  # 这一行是关键！
```

- **`resourcePolicy: "keep"`**：这个设置告诉 Helm，在执行 `helm uninstall harbor` 命令时，**不要删除** 由 Harbor 创建的 PVC。
- 由于 PVC 不会被删除，它所绑定的 PV 以及 PV 上的所有数据自然也就安全无虞。

当你重新使用 Helm 安装 Harbor 时，新的 Pod 会再次通过 PVC 挂载到原有的 PV 上，所有数据都完好如初。



### K8s 部署 Harbor 实战指南

本文档将指导您在 Kubernetes 集群中使用 Helm 部署 Harbor 私有镜像仓库，并配置 K8s 集群以拉取私有镜像。

#### 一、前提条件

- 一个正常运行的 Kubernetes 集群。
- 已安装并配置好 Helm v3。
- 集群中存在可用的 StorageClass，用于为 Harbor 提供持久化存储。



可以参考前面的有状态部署mysql，编写storagecalss和pv

```yaml
# --- 1. 定义 StorageClass ---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: harbor-sc           # 名字必须严格匹配
provisioner: kubernetes.io/no-provisioner  # 使用静态供给模式
volumeBindingMode: WaitForFirstConsumer    # 等待 Pod 调度后再绑定，避免节点不匹配问题

---
# --- 2. 创建 Registry 的 PV (对应配置中的 5Gi) ---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: harbor-registry-pv
spec:
  capacity:
    storage: 5Gi  
  accessModes:
    - ReadWriteOnce   # 如果是单节点部署用这个；多节点集群建议用 ReadWriteMany 并配合 NFS
  persistentVolumeReclaimPolicy: Retain
  storageClassName: harbor-sc  # 【关键】这里必须填 harbor-sc
  hostPath:
    path: /data/harbor/registry # 节点上的实际路径

---
# --- 3. 创建 Database 的 PV (对应配置中的 1Gi) ---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: harbor-database-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: harbor-sc  # 【关键】这里必须填 harbor-sc
  hostPath:
    path: /data/harbor/database # 节点上的实际路径
```



务必在**你的 K8s 节点服务器**上执行以下 Shell 命令，创建目录并赋予权限（否则 Harbor 的 Pod 会因为权限不足无法写入数据而启动失败）

```bash
# 1. 在服务器上创建目录
mkdir -p /data/harbor/registry
mkdir -p /data/harbor/database

# 2. 赋予权限 (777 是为了确保 Harbor 容器内用户能读写，生产环境可按需收紧)
chmod 777 /data/harbor/registry
chmod 777 /data/harbor/database

# 3. 应用配置到 K8s
kubectl apply -f harbor-storage.yaml
```

### **验证状态**

执行完上述步骤后，检查 PV 状态：

```
c
```



#### 二、部署 Harbor

我们将使用 Helm Chart 来部署 Harbor，这是推荐且最便捷的方式。

1. **添加 Harbor Helm 仓库**
   首先，将 Harbor 官方的 Helm 仓库添加到您的本地 Helm 客户端。

```bash
helm repo add harbor https://helm.goharbor.io
helm repo update
```

1. **创建 Harbor 命名空间**
   为 Harbor 创建一个独立的命名空间，以便于资源管理。

```bash
kubectl create namespace harbor
```

1. **准备配置文件**
   Harbor 的部署行为由 `values.yaml` 文件控制。您需要创建一个自定义的 `values.yaml` 文件来适配您的环境。

以下是关键配置的示例：

```yaml
# --- 1. 暴露方式与 HTTPS 配置 ---
expose:
  # 修改 1: 改为 nodePort (推荐用于无域名/HTTP环境) 或 clusterIP
  # 如果是 nodePort，可以通过 节点IP:端口 直接访问
  # 如果是 clusterIP，通常需要配合 Ingress Controller (需额外配置允许 HTTP)
  type: nodePort 
  
  tls:
    # 修改 2: 彻底关闭 HTTPS
    enabled: false

  # 如果使用 nodePort，可以在这里指定端口，不指定则随机分配
  nodePort:
    name: http
    port: 80
    nodePort: 30302 
externalURL: "http://47.115.225.81:30302"

# --- 2. 外部 Redis 配置 ---
externalRedis:
  # 修改 3: 填入外部 Redis 信息
  host: "10.43.6.245"
  port: 30379           
  
  # 如果有密码，填入密码；如果没有，留空
  password: "asd1234567-" 
  
  # 连接池配置 (可选，通常默认即可)
  core:
    poolSize: 100
  registry:
    poolSize: 100
  trivy:
    poolSize: 100

# --- 3. 禁用内置 Redis ---
redis:
  # 修改 4: 禁用内置的 Redis 组件
  enabled: false
  # 如果之前配置了内置 Redis 的持久化，这里可以全部注释掉以节省资源
  # persistence: ...

# --- 4. 其他原有配置 (保持不变) ---
harborAdminPassword: "asd1234567-"

persistence:
  enabled: true
  persistentVolumeClaim:
    registry:
      storageClass: "harbor-sc"
      size: 5Gi
    database:
      storageClass: "harbor-sc"
      size: 1Gi
```

下面是不使用外部redis的：

```yaml
# --- 1. 暴露方式与 HTTPS 配置 ---
expose:
  type: nodePort 
  
  tls:
    enabled: false

  nodePort:
    name: http
    port: 80
    nodePort: 30302 
externalURL: "http://47.115.225.81:30302"


# --- 2. 启用内置 Redis ---
redis:
  enabled: true  # 修改点：将 false 改为 true，启用内置 Redis
  # 可以使用默认的持久化配置，或根据需要进行调整
  # persistence:
  #   enabled: true
  #   storageClass: "harbor-sc"
  #   accessMode: ReadWriteOnce
  #   size: 1Gi

# --- 3. 其他原有配置 (保持不变) ---
harborAdminPassword: "asd1234567-"

persistence:
  enabled: true
  persistentVolumeClaim:
    registry:
      storageClass: "harbor-sc"
      size: 5Gi
    database:
      storageClass: "harbor-sc"
      size: 1Gi
```





**注意**：

- 请将 `harbor.example.com` 替换为您实际配置的域名。
- 请确保 `storageClass` 的值是您集群中真实存在的 StorageClass。
- 如果启用了 HTTPS (`tls.enabled: true`)，您需要提前创建好包含证书和私钥的 Secret。

```bash
# 创建 TLS Secret 示例
kubectl create secret tls harbor-tls \
  --cert=path/to/your/cert.crt \
  --key=path/to/your/private.key \
  -n harbor
```

1. **执行部署**
   使用准备好的 `values.yaml` 文件安装 Harbor。

```bash
helm install harbor harbor/harbor -n harbor -f harborvalues.yaml
```

1. **验证部署**
   等待所有 Pod 进入 `Running` 状态。

```bash
kubectl get pods -n harbor
```

检查 Ingress 是否已正确创建。

```bash
kubectl get ingress -n harbor
```





修改更新：

```bash
# 假设你的 release 名字叫 harbor
helm upgrade harbor harbor/harbor -n harbor -f harborvalues.yaml
```

卸载：

```bash
helm uninstall harbor -n harbor
```





踩过的大坑：

不使用http时，需要externalURL: "http://47.115.225.81:30302"，协议和端口都要，否则会出现登录不了的情况：{"errors":[{"code":"FORBIDDEN","message":"CSRF token invalid"}]}

admin

asd1234567-



目前还不知道为什么在配置文件里面写了端口是30302，但是每次helm upgrade后端口都变成30002，不知道我为什么，受影响的范围是harborvalues.yaml那里需要externalUrl端口需要对其，然后service每次upgrade后需要手动修改他



docker login 47.115.225.81:30302





```bash
 kubectl get svc -n harbor
NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
harbor-core         ClusterIP   10.43.234.90    <none>        80/TCP              41m
harbor-database     ClusterIP   10.43.34.51     <none>        5432/TCP            41m
harbor-jobservice   ClusterIP   10.43.242.76    <none>        80/TCP              41m
harbor-portal       ClusterIP   10.43.188.82    <none>        80/TCP              41m
harbor-redis        ClusterIP   10.43.154.211   <none>        6379/TCP            41m
harbor-registry     ClusterIP   10.43.189.245   <none>        5000/TCP,8080/TCP   41m
harbor-trivy        ClusterIP   10.43.219.251   <none>        8080/TCP            41m
http                NodePort    10.43.26.217    <none>        80:30302/TCP        41m

```





在docker推镜像时出现问题

决定使用https

生成证书

```bash
# 1. 创建目录
mkdir -p ~/harbor-cert && cd ~/harbor-cert

# 2. 生成私钥
openssl genrsa -out harbor.key 2048

# 3. 生成证书签名请求 (CSR)
# 注意：CN 后面必须填你的域名或 IP
openssl req -new -key harbor.key -out harbor.csr -subj "/C=CN/ST=Beijing/L=Beijing/O=Harbor/OU=IT/CN=47.115.225.81"

# 4. 生成自签名证书 (有效期 3650 天)
openssl x509 -req -days 3650 -in harbor.csr -signkey harbor.key -out harbor.crt

# 此时你会得到两个关键文件：harbor.key 和 harbor.crt

```

然后创建secret配置：

```bash
kubectl create secret tls harbor-tls \
  --cert=harbor.crt \
  --key=harbor.key \
  -n harbor
```

修改配置文件：
```yaml
# --- 1. 暴露方式与 HTTPS 配置 ---
expose:
  # 访问类型：nodePort
  type: nodePort 
  
  # --- 核心修改：开启 HTTPS ---
  tls:
    enabled: true
    # 证书来源：secret (对应 kubectl create secret tls harbor-tls ...)
    certSource: secret
    secret:
      # 这里必须填你刚才创建的 Secret 名字
      secretName: "harbor-tls"

  nodePort:
    name: harbor
    ports:
      # --- 核心修改：配置 HTTPS 端口 ---
      # 外部访问端口 (建议用 30443，避免和之前的 HTTP 端口冲突)
      https:
        port: 30443
        # 内部容器监听端口 (Harbor Nginx 的 HTTPS 标准端口是 8443)
        targetPort: 8443
        nodePort: 30443

# --- 核心修改：外部访问地址改为 HTTPS ---
externalURL: "https://47.115.225.81:30443"

# --- 2. 管理员密码配置 ---
harborAdminPassword: "asd1234567-"

# --- 3. 启用内置 Redis ---
redis:
  enabled: true

# --- 4. 持久化存储配置 ---
persistence:
  enabled: true
  # 这里指定了存储类，请确保你的 K8s 集群中存在名为 "harbor-sc" 的 StorageClass
  persistentVolumeClaim:
    registry:
      storageClass: "harbor-sc"
      size: 5Gi
    database:
      storageClass: "harbor-sc"
      size: 1Gi
    # 如果需要持久化其他组件（如 jobservice, redis, trivy），可以在这里继续添加

# --- 5. 其他默认配置 (保持默认即可) ---
# 日志级别
logLevel: info

# 数据库配置 (如果需要修改密码，请在这里改，否则保持默认)
database:
  password: "root123"
  # 最大连接数
  maxIdleConns: 25
  maxOpenConns: 50
```

### **客户端信任证书 (关键步骤)**

因为是**自签名证书**，Docker 客户端默认是不信任的。你有两种选择：

#### **方案 A：配置 Insecure Registry（简单粗暴，推荐测试用）**

在**执行 `docker login` 的机器**上，修改 `/etc/docker/daemon.json`：

json



```
{
  "insecure-registries": ["47.115.225.81:30443"]
}
```

然后重启 Docker：`systemctl restart docker`。

#### **方案 B：导入证书（生产环境推荐）**

如果你想让 Docker 信任这个证书，需要将 `harbor.crt` 复制到 Docker 的证书目录：

bash



```
# 1. 创建目录
sudo mkdir -p /etc/docker/certs.d/47.115.225.81:30443/

# 2. 复制证书 (注意文件名必须是 ca.crt)
# 假设你在第一步生成的证书还在当前目录下
sudo cp harbor.crt /etc/docker/certs.d/47.115.225.81:30443/ca.crt

# 3. 重启 Docker
sudo systemctl restart docker
```
