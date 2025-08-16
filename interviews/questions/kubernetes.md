# ☸️ Kubernetes面试题

> 适合：K8s工程师、SRE工程师、运维开发
> 难度：⭐⭐⭐⭐ (中级-高级)

## 📋 基础操作

### 1. Pod管理

#### 查看Node上的Pod
```bash
# 查看所有Pod及其所在节点
kubectl get pods -o wide

# 查看特定节点的Pod
kubectl get pods -o wide --field-selector spec.nodeName=node-1

# 查看节点详细信息
kubectl describe node node-1
```

#### Pod调度控制
```yaml
# 节点选择器
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
  - name: nginx
    image: nginx
  nodeSelector:
    disk: ssd

# 节点亲和性
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: kubernetes.io/e2e-az-name
            operator: In
            values:
            - e2e-az1
```

### 2. 健康检查

```yaml
# Liveness Probe - 存活探针
livenessProbe:
  httpGet:
    path: /health
    port: 80
  initialDelaySeconds: 30
  periodSeconds: 10

# Readiness Probe - 就绪探针  
readinessProbe:
  httpGet:
    path: /ready
    port: 80
  initialDelaySeconds: 5
  periodSeconds: 5
```

### 3. Service类型

```yaml
# ClusterIP（默认）
spec:
  type: ClusterIP
  ports:
  - port: 80
    targetPort: 80

# NodePort
spec:
  type: NodePort
  ports:
  - port: 80
    targetPort: 80
    nodePort: 30000

# Headless Service
spec:
  clusterIP: None
  selector:
    app: nginx
```

## 📋 高级管理

### 4. 大规模集群架构

```yaml
# 控制平面高可用
- 3-5个master节点
- etcd集群部署（5个节点）
- API Server水平扩展
- 负载均衡器

# 网络方案
- Calico BGP模式：适合大规模集群
- Flannel VXLAN：简单但性能一般

# 存储方案
- 分布式存储：Ceph、GlusterFS
- 云原生存储：Longhorn、Rook
```

### 5. 多租户资源隔离

```yaml
# 命名空间隔离
apiVersion: v1
kind: ResourceQuota
metadata:
  name: compute-quota
  namespace: tenant-a
spec:
  hard:
    requests.cpu: "4"
    requests.memory: 8Gi
    limits.cpu: "8"
    limits.memory: 16Gi

# 网络策略
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all
spec:
  podSelector: {}
  policyTypes:
  - Ingress
  - Egress
```

### 6. 升级策略

```yaml
# 滚动更新
strategy:
  type: RollingUpdate
  rollingUpdate:
    maxSurge: 25%
    maxUnavailable: 25%

# 升级控制
kubectl rollout status deployment nginx-deployment
kubectl rollout pause deployment nginx-deployment
kubectl rollout undo deployment nginx-deployment
```

## 📋 故障排查

### 7. 常见问题诊断

```bash
# Pod故障排查
kubectl get pods
kubectl describe pod pod-name
kubectl logs pod-name

# 资源使用检查
kubectl top pods
kubectl top nodes

# 网络连接测试
kubectl exec -it pod-name -- /bin/sh
ping service-name
curl http://service-name:port

# 存储检查
kubectl get pv,pvc
kubectl describe pvc pvc-name
```

### 8. 性能优化

- 调整kubelet参数
- 优化调度器配置
- 网络插件调优
- 存储性能优化

## 🔗 相关资源

- [Kubernetes官方文档](https://kubernetes.io/docs/)
- [K8s最佳实践](https://kubernetes.io/docs/setup/best-practices/)

## 📝 复习要点

1. **掌握K8s基础命令和操作**
2. **理解Pod调度和节点选择机制**
3. **学会健康检查和故障排查方法**
4. **熟悉大规模集群管理**
