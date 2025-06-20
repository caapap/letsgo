# Kubernetes核心原理详解

> 本文档是《Go运维开发工程师认证考试复习指南》的K8s核心原理详解补充部分

## 1. Service如何控制Ingress网络流量的机制与原理

### 核心架构关系
Service和Ingress在K8s网络架构中处于不同抽象层次：
- **Service**: L4层服务抽象，提供稳定的网络端点和负载均衡
- **Ingress**: L7层路由规则，定义外部访问的HTTP/HTTPS流量分发策略

### 技术实现原理

#### 1. 资源关联机制
```yaml
# Service定义 - 提供稳定的内部访问端点
apiVersion: v1
kind: Service
metadata:
  name: backend-service
spec:
  selector:
    app: backend
  ports:
  - name: http
    port: 80
    targetPort: 8080
    protocol: TCP
  type: ClusterIP

---
# Ingress定义 - 通过service.name建立关联
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: backend-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  ingressClassName: nginx
  rules:
  - host: api.example.com
    http:
      paths:
      - path: /api
        pathType: Prefix
        backend:
          service:
            name: backend-service  # 关键：通过名称引用Service
            port:
              number: 80
```

#### 2. 流量转发链路
```
外部请求 → LoadBalancer/NodePort → Ingress Controller → Service → Endpoints → Pod
```

#### 3. 底层实现机制
- **Ingress Controller**监听Ingress资源变更，动态生成负载均衡配置
- **Service**通过EndpointSlice维护后端Pod列表
- **kube-proxy**维护iptables/ipvs规则，实现Service到Pod的负载均衡
- **DNS解析**：Service在集群内部提供稳定的DNS记录

### 关键技术细节
1. **服务发现**: Ingress通过Service名称进行服务发现，而非直接访问Pod IP
2. **会话保持**: 可通过Service的sessionAffinity和Ingress的sticky session实现
3. **健康检查**: Service的readinessProbe确保只有健康的Pod接收流量

---

## 2. 三Master节点高可用集群的实现原理

### 核心组件分布式架构

#### etcd集群（数据层高可用）
```bash
# 三节点etcd集群配置示例
# Master1: etcd member 1
# Master2: etcd member 2  
# Master3: etcd member 3
```

**Raft协议保证数据一致性**：
- **Leader选举**: 任意时刻只有一个Leader处理写请求
- **日志复制**: Leader将操作日志复制到Follower节点
- **一致性保证**: 需要超过半数节点(≥2)确认才能提交写操作

#### API Server（接入层高可用）
- **无状态设计**: 多个API Server实例可并行运行
- **负载均衡**: 通过HAProxy/Nginx实现请求分发
- **数据一致性**: 所有实例共享同一个etcd集群

#### Controller Manager & Scheduler（控制层高可用）
```yaml
# Leader Election配置
spec:
  containers:
  - command:
    - kube-controller-manager
    - --leader-elect=true
    - --leader-elect-lease-duration=15s
    - --leader-elect-renew-deadline=10s
    - --leader-elect-retry-period=2s
    - --leader-elect-resource-lock=leases
```

**Leader Election机制**：
- 基于etcd的分布式锁实现
- 同时只有一个实例处于Active状态
- Leader故障时自动触发重新选举

### 故障容忍能力
- **etcd**: 可容忍(n-1)/2个节点故障（3节点集群可容忍1个故障）
- **API Server**: 任意数量节点故障，只要有一个节点正常即可
- **Controller/Scheduler**: 自动故障转移，RTO < 30秒

---

## 3. Master节点故障后的集群恢复机制

### etcd故障恢复原理

#### 单节点故障（保持仲裁）
**自动恢复流程**：
1. **故障检测**: 其他节点通过心跳检测到故障节点
2. **Leader重选**: 如果故障节点是Leader，触发新的Leader选举
3. **服务继续**: 集群在剩余节点上继续提供服务
4. **节点恢复**: 故障节点重启后自动同步数据并重新加入集群

#### 多节点故障（失去仲裁）
**手动恢复步骤**：
```bash
# 1. 停止所有etcd服务
systemctl stop etcd

# 2. 从最新快照恢复数据
ETCDCTL_API=3 etcdctl snapshot restore /backup/snapshot.db \
  --name=master1 \
  --initial-cluster="master1=https://10.0.1.10:2380" \
  --initial-cluster-token="etcd-cluster-token" \
  --initial-advertise-peer-urls="https://10.0.1.10:2380" \
  --data-dir="/var/lib/etcd-restore"

# 3. 更新etcd配置，重建单节点集群
# 4. 启动第一个etcd节点
systemctl start etcd

# 5. 逐步添加其他节点
etcdctl member add master2 --peer-urls="https://10.0.1.11:2380"
```

### Kubernetes组件恢复顺序

#### 1. etcd恢复（数据层）
```bash
# 验证etcd集群健康状态
ETCDCTL_API=3 etcdctl endpoint health --cluster
ETCDCTL_API=3 etcdctl member list
```

#### 2. API Server恢复（接入层）
```bash
# 检查API Server配置
cat /etc/kubernetes/manifests/kube-apiserver.yaml

# 验证API Server连接etcd
kubectl get componentstatuses
```

#### 3. Controller Manager & Scheduler恢复（控制层）
```bash
# 检查Leader Election状态
kubectl get lease -n kube-system
kubectl describe lease kube-controller-manager -n kube-system
```

#### 4. 节点组件恢复（数据层）
```bash
# 重启kubelet和kube-proxy
systemctl restart kubelet
kubectl delete pod -n kube-system -l k8s-app=kube-proxy
```

### 关键配置文件
```bash
# etcd配置
/etc/etcd/etcd.conf
/etc/kubernetes/pki/etcd/

# Kubernetes组件配置
/etc/kubernetes/manifests/
/etc/kubernetes/admin.conf
/var/lib/kubelet/config.yaml
```

---

## 4. NodePort服务无法访问的系统化排查方法

### 分层诊断方法论

#### Layer 1: Pod层面检查
```bash
# 1. 验证Pod运行状态
kubectl get pods -o wide
kubectl describe pod <pod-name>

# 2. 测试Pod内部服务
kubectl exec -it <pod-name> -- curl localhost:8080
kubectl exec -it <pod-name> -- netstat -tlnp
```

#### Layer 2: Service层面检查
```bash
# 1. 检查Service配置
kubectl get svc -o wide
kubectl describe svc <service-name>

# 2. 验证Endpoints
kubectl get endpoints <service-name>
kubectl get endpointslices -l kubernetes.io/service-name=<service-name>

# 3. 测试Service内部访问
kubectl run debug --image=busybox --rm -it -- sh
# 在debug pod中: wget -qO- <service-name>:<port>
```

#### Layer 3: NodePort层面检查
```bash
# 1. 确认NodePort端口分配
kubectl get svc <service-name> -o jsonpath='{.spec.ports[0].nodePort}'

# 2. 检查端口监听状态
netstat -tlnp | grep :<nodeport>
ss -tlnp | grep :<nodeport>

# 3. 验证节点内部访问
curl localhost:<nodeport>
```

#### Layer 4: 网络层面检查
```bash
# 1. 检查防火墙规则
systemctl status firewalld
iptables -L -n | grep <nodeport>
firewall-cmd --list-ports

# 2. 检查kube-proxy规则
iptables -t nat -L KUBE-NODEPORTS -n
ipvsadm -L -n  # 如果使用ipvs模式

# 3. 检查CNI网络
kubectl get pods -n kube-system -l app=<cni-plugin>
ip route show
```

### 常见问题模式与解决方案

#### 1. 标签选择器不匹配
```bash
# 问题诊断
kubectl get pods --show-labels
kubectl get svc <service-name> -o yaml | grep selector

# 解决方案：确保Service selector与Pod labels完全匹配
```

#### 2. kube-proxy组件异常
```bash
# 问题诊断
kubectl get pods -n kube-system -l k8s-app=kube-proxy
kubectl logs -n kube-system -l k8s-app=kube-proxy

# 解决方案：重启kube-proxy
kubectl delete pod -n kube-system -l k8s-app=kube-proxy
```

#### 3. 防火墙策略阻挡
```bash
# 解决方案：开放NodePort端口范围
firewall-cmd --permanent --add-port=30000-32767/tcp
firewall-cmd --reload

# 或者添加iptables规则
iptables -I INPUT -p tcp --dport 30000:32767 -j ACCEPT
```

#### 4. 网络策略限制
```bash
# 检查NetworkPolicy
kubectl get networkpolicy --all-namespaces
kubectl describe networkpolicy <policy-name>
```

### 排查工具集
```bash
# 网络连通性测试
telnet <node-ip> <nodeport>
nc -zv <node-ip> <nodeport>

# 流量抓包分析
tcpdump -i any port <nodeport>

# 集群网络诊断
kubectl run netshoot --rm -i --tty --image nicolaka/netshoot -- /bin/bash
```

---

## 面试回答技巧

### 回答结构建议
1. **原理概述**: 先说整体架构和核心原理
2. **技术细节**: 深入关键实现机制
3. **实际经验**: 结合具体配置和命令
4. **故障处理**: 展示问题排查和解决能力

### 加分要点
- 能够画出架构图说明组件关系
- 熟悉相关配置文件和关键参数
- 具备实际的故障排查经验
- 了解不同场景下的最佳实践 