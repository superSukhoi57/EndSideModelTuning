在 Kubernetes (k8s) 中使用 StatefulSet 部署 MySQL 是一个常见的实践，因为 StatefulSet 能够为有状态应用提供稳定的网络标识和持久化存储。



# 部署流程概览

1. **创建无头服务 (Headless Service)**：为 StatefulSet 中的每个 Pod 提供稳定的网络标识。
2. **创建持久化存储 (Persistent Volumes)**：为 MySQL 数据提供持久化存储。
3. **部署 MySQL StatefulSet**：定义并部署 MySQL 实例。
4. **验证部署**：检查 Pod 和服务状态。

------

## 详细步骤

### 1. 创建无头服务 (Headless Service)

首先，需要创建一个无头服务，它的作用是为 StatefulSet 管理的 Pod 提供稳定的 DNS 记录。每个 Pod 将获得一个唯一的 DNS 名称，格式为 `<pod-name>.<service-name>.<namespace>.svc.cluster.local`。

创建一个名为 `mysql-service.yaml` 的文件：

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql
  labels:
    app: mysql
spec:
  ports:
  - port: 3306
    name: mysql
  clusterIP: None # 关键配置，定义无头服务
  selector:
    app: mysql
```

**关键点**：`clusterIP: None` 是定义无头服务的核心。

应用该配置：

```bash
kubectl apply -f mysql-service.yaml
```

### 2. 创建持久化存储 (Persistent Volumes)

StatefulSet 通过 `volumeClaimTemplates` 来为每个 Pod 动态创建持久化卷声明 (PVC)。这需要一个预先存在的 StorageClass 来提供存储。

这里我们创建一个简单的 StorageClass 用于演示，它使用 `hostPath`（仅适用于单节点测试环境）。在生产环境中，应使用云厂商或集群管理员提供的 StorageClass。

创建一个名为 `mysql-sc.yaml` 的文件：

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: manual
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
```

应用该配置：

```bash
kubectl apply -f mysql-sc.yaml
```

**这段代码创建的是一个 StorageClass（存储类），而不是 PV（PersistentVolume）**

提供的 YAML 定义了一个 **StorageClass**，名字叫 `manual`。

- **它的角色**：它像是一个“**模版**”或者“**类别标签**”。它告诉 Kubernetes：“凡是申请 `manual` 这种类型的存储，都要按照我规定的规则来办。”
- **关键配置解读**：
  - `provisioner: kubernetes.io/no-provisioner`：这句话是核心。它的意思是 **“Kubernetes，你不要自动帮我创建存储（PV），我自己会手动创建。”**
  - `volumeBindingMode: WaitForFirstConsumer`：意思是“等到有 Pod 真正要用这个存储时，再进行绑定”。

#### 创建真正的pv

//mysql-pv.yaml

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: mysql-pv-0  # PV 的名字
spec:
  capacity:
    storage: 10Gi   # 容量必须 >= PVC 请求的大小
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: manual  # 关键：必须匹配你的 StorageClass 名字
  hostPath:
    path: /data/pvcmnt # 节点上的实际路径
```

那里的path就是机器上的实际路径

查看结果

```
root@iZf8z0p5x4accyquvjovsnZ:/etc/rancher/k3s# kubectl get pv
NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                                      STORAGECLASS   VOLUMEATTRIBUTESCLASS   REASON   AGE
mysql-pv-0   10Gi       RWO            Retain           Bound    default/mysql-persistent-storage-mysql-0   manual         <unset>                          2m32s
root@iZf8z0p5x4accyquvjovsnZ:/etc/rancher/k3s# kubectl get pvc
NAME                               STATUS   VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   VOLUMEATTRIBUTESCLASS   AGE
mysql-persistent-storage-mysql-0   Bound    mysql-pv-0   10Gi       RWO            manual         <unset>                 2m40s

```

你会看到 PV 的状态变成 `Bound`，PVC 的状态也会变成 `Bound`。

| 概念 | 全称                  | 角色   | 谁负责创建？                      | 比喻                                            |
| ---- | --------------------- | ------ | --------------------------------- | ----------------------------------------------- |
| PV   | PersistentVolume      | 供应方 | 运维/管理员 (或者你手动创建)      | 房子 (真实的物理资源，有具体的地址、面积)       |
| PVC  | PersistentVolumeClaim | 消费方 | 开发者/用户 (在 StatefulSet 里写) | 租房订单 (我的需求：我要住市中心，至少 10 平米) |

- **PV (`mysql-pv-0`)**：是你手动创建的那块真实的硬盘目录。
- **PVC (`mysql-persistent-storage-mysql-0`)**：是 MySQL Pod 手里拿的“钥匙”。
- **Pod**：拿着这把钥匙，打开了那扇门，开始往里面存数据。



### 3. 部署 MySQL StatefulSet

接下来，创建 StatefulSet 资源。这个配置会创建 3 个 MySQL 实例，并为每个实例分配独立的持久化存储。

创建一个名为 `mysql-statefulset.yaml` 的文件：

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
spec:
  serviceName: "mysql" # 必须与之前创建的无头服务名称一致
  replicas: 1
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - name: mysql
        image: mysql:8.0
        ports:
        - containerPort: 3306
          name: mysql
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: "asd1234567-" # 请替换为强密码
        volumeMounts:
        - name: mysql-persistent-storage
          mountPath: /var/lib/mysql # MySQL 数据目录
  volumeClaimTemplates:
  - metadata:
      name: mysql-persistent-storage
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: manual # 引用上一步创建的 StorageClass
      resources:
        requests:
          storage: 5Gi # 为每个 Pod 申请 10Gi 的存储空间
```

**关键点**：

- `serviceName`: 必须指向之前创建的无头服务。
- `volumeClaimTemplates`: 这是 StatefulSet 的核心，它为每个 Pod 自动创建一个 PVC，确保数据的持久性和独立性。

应用该配置：

```bash
kubectl apply -f mysql-statefulset.yaml
```

4. 验证部署

部署完成后，可以通过以下命令检查资源状态。

- **查看 StatefulSet 状态**：

  ```bash
  kubectl get statefulset mysql
  ```

- **查看 Pod 状态**：

  ```bash
  kubectl get pods -l app=mysql
  ```

  你应该能看到 `mysql-0`, `mysql-1`, `mysql-2` 三个 Pod，并且它们的状态都是 `Running`。StatefulSet 会按顺序（0, 1, 2...）启动 Pod。

- **查看 PVC 状态**：

  ```bash
  kubectl get pv
  kubectl get pvc
  ```

  你会看到与 Pod 对应的三个 PVC，例如 `mysql-persistent-storage-mysql-0`，并且它们的状态是 `Bound`。



### 暴露服务

我们可以在statefulset查看这个服务的yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: '2026-04-22T08:03:11Z'
  labels:
    app: mysql
  managedFields:
    - apiVersion: v1
      fieldsType: FieldsV1
      fieldsV1:
        f:metadata:
          f:labels:
            .: {}
            f:app: {}
        f:spec:
          f:clusterIP: {}
          f:internalTrafficPolicy: {}
          f:ports:
            .: {}
            k:{"port":3306,"protocol":"TCP"}:
              .: {}
              f:name: {}
              f:port: {}
              f:protocol: {}
              f:targetPort: {}
          f:selector: {}
          f:sessionAffinity: {}
          f:type: {}
      manager: dashboard-api
      operation: Update
      time: '2026-04-22T08:03:11Z'
  name: mysql
  namespace: default
  resourceVersion: '7897'
  uid: c5220651-5420-40fe-ba70-72ad3df9b05a
spec:
  clusterIP: None
  clusterIPs:
    - None
  internalTrafficPolicy: Cluster
  ipFamilies:
    - IPv4
  ipFamilyPolicy: SingleStack
  ports:
    - name: mysql
      port: 3306
      protocol: TCP
      targetPort: 3306
  selector:
    app: mysql
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```



#### service类型

| Service 类型      | 核心特性                                                     | 访问范围    | 典型应用场景                                                 |
| ----------------- | ------------------------------------------------------------ | ----------- | ------------------------------------------------------------ |
| ClusterIP         | 默认类型。分配一个仅集群内部可见的虚拟 IP。                  | 集群内部    | 微服务之间的内部通信（如前端调后端），不对外暴露。           |
| NodePort          | 在每个节点上开放一个固定静态端口（NodePort，范围通常是 30000-32767）。 | 外部 + 内部 | 开发/测试环境，或没有云负载均衡器时，通过 `<节点IP>:<端口>` 访问。 |
| LoadBalancer      | 请求云厂商创建一个外部负载均衡器并分配公网 IP。              | 互联网      | 生产环境，需要向公网用户提供服务（如 Web 应用入口）。        |
| ExternalName      | 将服务映射到外部域名（DNS CNAME 记录）。                     | 集群内部    | 让集群内的应用通过统一的服务名访问集群外的数据库或 API。     |
| Headless *(配置)* | 设置 `clusterIP: None`，不分配虚拟 IP。                      | 集群内部    | 有状态应用（如 MySQL 主从、Zookeeper），客户端需要直接连接特定 Pod IP。 |

快速选择：

- **只在内部用？** 选 `ClusterIP`。
- **想从公司外网测试？** 选 `NodePort`。
- **要给用户通过公网访问？** 选 `LoadBalancer`。
- **要连外部的阿里云 RDS？** 选 `ExternalName`。

所以需要修改：

```yaml
  ports:
    - name: mysql
      port: 3306
      protocol: TCP
      targetPort: 3306
      nodePort: 30306  # 新增这一行，指定外部访问端口
  selector:
    app: mysql
  sessionAffinity: None
  type: NodePort
status:
  loadBalancer: {}
```



