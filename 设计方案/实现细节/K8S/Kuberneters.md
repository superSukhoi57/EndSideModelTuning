以API编程的方式管理安排各个容器的引擎

[kubernetes(k8s)课程.pdf](D:/the_files_at/Assignment/笔记资料/尚硅谷Kubernetes（k8s）新版/笔记/课程笔记/kubernetes(k8s)课程.pdf)



K8s 功能: 

（1）自动装箱： 基于容器对应用运行环境的资源配置要求自动部署应用容器 （2）自我修复(自愈能力)： 当容器失败时，会对容器进行重启 当所部署的 Node 节点有问题时，会对容器进行重新部署和重新调度 当容器未通过监控检查时，会关闭此容器直到容器正常运行时，才会对外提供服务 （3）水平扩展 ：通过简单的命令、用户 UI 界面或基于 CPU 等资源使用情况，对应用容器进行规模扩大 或规模剪裁 （3）服务发现 用户不需使用额外的服务发现机制，就能够基于 Kubernetes 自身能力实现服务发现和 负载均衡 （4）滚动更新： 可以根据应用的变化，对应用容器运行的应用，进行一次性或批量式更新 （5）版本回退： 可以根据应用部署情况，对应用容器运行的应用，进行历史版本即时回退 （6）密钥和配置管理 ：在不需要重新构建镜像的情况下，可以部署和更新密钥和应用配置，类似热部署。 （7）存储编排： 自动实现存储系统挂载及应用，特别对有状态应用实现数据持久化非常重要 存储系统可以来自于本地目录、网络存储(NFS、Gluster、Ceph 等)、公共云存储服务 （8）批处理 ：提供一次性任务，定时任务；满足批量数据处理和分析的场景



架构分为主控系节点master和工作节点



核心概念：

1、Pod：最小部署单元，一组容器的集合，共享网络（说的是Pod里的某个容器监听80，那么这个Pod里的所有容器也会获得80），生命周期是短暂的

2、controller：确保预期的Pod副本数量，有/无状态应用部署，确保所有的node运行同一个pod，执行一次性任务和定时任务

3、service：定义了一组pod的访问规则



kubeadm搭建集群：

这个工具能通 过两条指令完成一个 kubernetes 集群的部署： 第一、创建一个 Master 节点 kubeadm init 

第二， 将 Node 节点加入到当前集群中 $ kubeadm join 







# K3S安装

[(37 封私信 / 80 条消息) 轻量级Kubernetes，在Linux上运行K3s - 知乎](https://zhuanlan.zhihu.com/p/1994886076926542836)

主要关注他将可执行文件放哪里了。

> r.cn/k3s/k3s-install.sh: No such file or directory
> root@iZf8z0p5x4accyquvjovsnZ:~/k3s# curl -sfL https://rancher-mirror.rancher.cn/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn sh -
> [INFO]  Finding release for channel stable
> [INFO]  Using v1.34.6+k3s1 as release
> [INFO]  Downloading hash rancher-mirror.rancher.cn/k3s/v1.34.6-k3s1/sha256sum-amd64.txt
> [INFO]  Downloading binary rancher-mirror.rancher.cn/k3s/v1.34.6-k3s1/k3s
> [INFO]  Verifying binary download
> [INFO]  Installing k3s to /usr/local/bin/k3s
> [INFO]  Skipping installation of SELinux RPM
> [INFO]  Creating /usr/local/bin/kubectl symlink to k3s
> [INFO]  Creating /usr/local/bin/crictl symlink to k3s
> [INFO]  Creating /usr/local/bin/ctr symlink to k3s
> [INFO]  Creating killall script /usr/local/bin/k3s-killall.sh
> [INFO]  Creating uninstall script /usr/local/bin/k3s-uninstall.sh
> [INFO]  env: Creating environment file /etc/systemd/system/k3s.service.env
> [INFO]  systemd: Creating service file /etc/systemd/system/k3s.service
> [INFO]  systemd: Enabling k3s unit
> Created symlink /etc/systemd/system/multi-user.target.wants/k3s.service → /etc/systemd/system/k3s.service.
> [INFO]  systemd: Starting k3s

之后就可以运行下面的命令检查了：

```
# 检查K3s服务状态: active（running）
systemctl status k3s
# 查看集群节点信息: Ready control-plane
kubectl get nodes
# 查看所有运行中的Pod
kubectl get pods -A
```

### **解决网络问题**

K3s访问的镜像仓库默认为DockerHub，在云服务的网络环境下会遇到镜像拉取失败问题。可配置镜像加速器（换源）：

```bash
# 创建镜像加速配置
sudo mkdir -p /etc/rancher/k3s
sudo cat > /etc/rancher/k3s/registries.yaml << EOF
mirrors:
  docker.io:
    endpoint:
      - "https://registry.cn-hangzhou.aliyuncs.com/"
      - "https://mirror.ccs.tencentyun.com"
  quay.io:
    endpoint:
      - "https://quay.tencentcloudcr.com/"
  registry.k8s.io:
    endpoint:
      - "https://registry.aliyuncs.com/v2/google_containers"
  gcr.io:
    endpoint:
      - "https://gcr.m.daocloud.io/"
  k8s.gcr.io:
    endpoint:
      - "https://registry.aliyuncs.com/google_containers"
  ghcr.io:
    endpoint:
      - "https://ghcr.m.daocloud.io/"
EOF

# 重启K3s服务使配置生效
sudo systemctl restart k3s
```





查看状态：

```
kubectl get pods -A
NAMESPACE     NAME                                      READY   STATUS         RESTARTS   AGE
kube-system   coredns-76c974cb66-rwkw9                  1/1     Running        0          15m
kube-system   helm-install-traefik-2rl2d                0/1     ErrImagePull   0          15m
kube-system   helm-install-traefik-crd-hdrcn            0/1     ErrImagePull   0          15m
kube-system   local-path-provisioner-8686667995-cr9z7   1/1     Running        0          15m
kube-system   metrics-server-c8774f4f4-wfvqx            1/1     Running        0          15m

```

查看pod里面到底发生了什么：

```
 kubectl describe pod -n kube-system helm-install-traefik-2rl2d
```

在容器acr里找到镜像加速：https://3cuahd3m.mirror.aliyuncs.com

我看到的现象是还是用的旧的配置

通过**systemctl restart k3s**  重启后没删掉

```bash
# 例如删除 helm-install 相关的 pod，它会重建
kubectl delete pod -n kube-system -l job-name=helm-install-traefik
```

**正常的现象：**

```
 sudo kubectl get pods -A
NAMESPACE     NAME                                      READY   STATUS      RESTARTS   AGE
kube-system   coredns-76c974cb66-rwkw9                  1/1     Running     0          46m
kube-system   helm-install-traefik-crd-hdrcn            0/1     Completed   0          46m
kube-system   local-path-provisioner-8686667995-cr9z7   1/1     Running     0          46m
kube-system   metrics-server-c8774f4f4-wfvqx            1/1     Running     0          46m
kube-system   svclb-traefik-bf5c8045-vq967              2/2     Running     0          13m
kube-system   traefik-c5c8bf4ff-4gds2                   1/1     Running     0          13m
```

看见这个就是正常的状态，Completed个 Pod 是由 **Job** 控制器管理的，它的任务性质是**“一次性任务”**，而不是像 Nginx 或 MySQL 那样需要一直运行的服务。

**安装可视化工具：**

[安装Kubernetes Dashboard与Helm-开发者社区-阿里云](https://developer.aliyun.com/article/1572543)这里使用helm安装

[(37 封私信 / 80 条消息) Kubernetes：Dashboard 安装 - 知乎](https://zhuanlan.zhihu.com/p/1958237810021631421)包含离线安装

创建临时token：` kubectl -n kubernetes-dashboard create token admin-user`

## RBAC

//dashboard-user.yaml

```yaml
# 1. 创建 ServiceAccount
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
---
# 2. 绑定集群管理员权限
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: admin-user
    namespace: kubernetes-dashboard
```

这个是一个一步到位的配置在 Kubernetes 的 YAML 文件中，`---` 是**分隔符**。它的作用是告诉 Kubernetes：“这里是一个资源的结束，下面紧接着是另一个新资源的开始”。
因此，你可以把 `ServiceAccount` 和 `ClusterRoleBinding` 写在同一个文件里，通过 `kubectl apply -f` 一次性全部创建。

然后应用

```bash
kubectl apply -f dashboard-user.yaml
```



## 长期token

//dashboard-user-secret.yaml

```yaml
 apiVersion: v1
 kind: Secret
 metadata:
   name: admin-user
   namespace: kubernetes-dashboard
   annotations:
     kubernetes.io/service-account.name: "admin-user"   
 type: kubernetes.io/service-account-token  
```

然后执行：`kubectl apply -f dashboard-user-secret.yaml`

获取长期token

```bash
kubectl get secret admin-user -n kubernetes-dashboard -o jsonpath="{.data.token}" | base64 -d
```



获取到的token：

```bash
eyJhbGciOiJSUzI1NiIsImtpZCI6IkRfdW13eXZmaG04TTRsblF2bWRVZVNBVHVaMDJZMjJwSV9faDhQMlVFOXMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi11c2VyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJhZmU1M2MzYS02N2Q1LTQ4OWItOTc1Yi05Y2M5NTJlMTc3MDQiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZXJuZXRlcy1kYXNoYm9hcmQ6YWRtaW4tdXNlciJ9.Qc169vjLtGr0awtRanl5Rd80th12AhJCOYSM-sdHL0KMy2_QrrwjpPi_AXAMR28aDLE3FJKNGweMjrfL6--9SqQofomJb0F3mXbJjIyF0eZCfWdabX1KVIYuVlXqu9zTno_QyF3iPBc5SFIXVnLhqiMAicAHviWsWFr4RlgLee3yIqxPmyMY3wR5FxEWIp_fBDy8agPtj792lZNl3ZsdauqsvBAmwnESFbqviMqCO7t7AAeg7rAFt0X1SaijLODtjoNgazTGvK9nDSzAGARO2P6YhB1qe71L3nPB16m03RrXH0poS7LG8MKwLvONY2fQGr2yagnmIFaF9FLwA4XaIA
```

