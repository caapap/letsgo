# Kubernetes面试核心问题专业回答

> 本文档是《Go运维开发工程师认证考试复习指南》的K8s高级面试题补充部分

## 1. Service控制Ingress网络的过程和原理

### 核心机制
Service和Ingress是K8s网络的两个不同抽象层：
- **Service**: 四层负载均衡，提供稳定的集群内访问端点
- **Ingress**: 七层路由规则，定义外部HTTP/HTTPS访问策略

### 工作原理
```
外部流量 → Ingress Controller → Service → Endpoints → Pod
```

**关键实现**：
1. Ingress通过`backend.service.name`字段引用Service
2. Ingress Controller监听Ingress资源，动态生成Nginx/HAProxy配置
3. Service通过EndpointSlice维护健康的Pod列表
4. kube-proxy维护iptables/ipvs规则实现负载均衡

**技术细节**：
- Service提供服务发现和DNS解析
- Ingress Controller作为反向代理处理L7路由
- 通过readinessProbe确保流量只转发到健康Pod

---

## 2. 三Master高可用实现原理

### 核心架构
每个Master节点运行：etcd + API Server + Controller Manager + Scheduler

### 高可用机制

#### etcd集群（Raft协议）
- **一致性保证**: 写操作需要超过半数节点确认
- **Leader选举**: 任意时刻只有一个Leader处理写请求
- **故障容忍**: 3节点集群可容忍1个节点故障

#### API Server（无状态）
- 多实例并行运行，通过负载均衡器分发请求
- 共享同一个etcd集群，保证数据一致性

#### Controller Manager & Scheduler（Leader Election）
```bash
--leader-elect=true
--leader-elect-lease-duration=15s
```
- 基于etcd分布式锁实现Leader选举
- 同时只有一个实例Active，其他实例Standby
- Leader故障时自动重新选举，RTO < 30秒

---

## 3. Master宕机后的恢复机制

### etcd恢复原理

#### 单节点故障（保持仲裁）
**自动恢复**：
1. 其他节点检测到故障，重新选举Leader
2. 集群继续在剩余节点上提供服务
3. 故障节点恢复后自动同步数据重新加入

#### 多节点故障（失去仲裁）
**手动恢复步骤**：
```bash
# 1. 从备份恢复etcd数据
etcdctl snapshot restore /backup/snapshot.db \
  --data-dir=/var/lib/etcd-restore

# 2. 重建单节点etcd集群
# 3. 逐步添加其他Master节点
etcdctl member add master2 --peer-urls="https://ip:2380"
```

### 组件恢复顺序
1. **etcd恢复** → 验证集群健康状态
2. **API Server启动** → 检查与etcd连接
3. **Controller Manager/Scheduler** → 验证Leader Election
4. **节点组件** → 重启kubelet和kube-proxy

### 关键配置文件
```bash
/etc/etcd/etcd.conf
/etc/kubernetes/manifests/kube-apiserver.yaml
/etc/kubernetes/admin.conf
```

---

## 4. NodePort访问不通的排查方法

### 系统化排查流程

#### 1. Pod层检查
```bash
kubectl get pods -o wide
kubectl exec -it pod -- curl localhost:8080
```

#### 2. Service层检查
```bash
kubectl describe svc service-name
kubectl get endpoints service-name
```
**常见问题**: 标签选择器不匹配，Endpoints为空

#### 3. NodePort层检查
```bash
kubectl get svc -o wide  # 确认NodePort端口
netstat -tlnp | grep :30xxx  # 检查端口监听
curl localhost:30xxx  # 节点内部测试
```

#### 4. 网络层检查
```bash
# 防火墙检查
firewall-cmd --list-ports
iptables -L -n

# kube-proxy规则检查
iptables -t nat -L KUBE-NODEPORTS
kubectl logs -n kube-system -l k8s-app=kube-proxy
```

### 常见问题及解决方案

1. **防火墙阻挡**
```bash
firewall-cmd --add-port=30000-32767/tcp --permanent
```

2. **kube-proxy异常**
```bash
kubectl delete pod -n kube-system -l k8s-app=kube-proxy
```

3. **Service配置错误**
- 检查selector与Pod labels匹配
- 确认targetPort配置正确

4. **CNI网络问题**
```bash
kubectl get pods -n kube-system  # 检查网络插件状态
```

### 排查工具
```bash
telnet node-ip nodeport
tcpdump -i any port nodeport
kubectl run debug --image=busybox --rm -it -- sh
```

---

## 面试表达技巧

### 回答结构
1. **概念定义** → 说明核心概念和作用
2. **实现原理** → 解释底层技术机制  
3. **关键配置** → 展示具体配置和命令
4. **故障处理** → 体现实际操作经验

### 技术亮点
- 熟悉Raft协议、Leader Election等分布式算法
- 掌握iptables/ipvs、CNI等网络技术
- 具备系统化的故障排查方法论
- 了解K8s各组件的配置文件和关键参数

### 进阶话题准备
- etcd性能优化和监控
- Ingress Controller的选型和调优
- 网络策略和安全加固
- 集群升级和备份恢复策略 