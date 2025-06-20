# 🚀 核心知识点速查手册

> 考试前15分钟快速复习，涵盖所有高频考点

## 📋 目录
- [Kubernetes核心问题](#kubernetes核心问题)
- [Go语言要点](#go语言要点)
- [分布式系统](#分布式系统)
- [中间件技术](#中间件技术)
- [算法必备](#算法必备)

---

## Kubernetes核心问题

### 🔍 Image Pending 情况分析
**镜像拉取问题**
- ❌ 镜像不存在或标签错误
- ❌ 网络连接问题（仓库不可达）
- ❌ 认证失败（私有仓库凭据错误）
- ❌ 镜像过大导致拉取超时

**资源调度问题**
- ❌ 节点资源不足（CPU/内存/存储）
- ❌ nodeSelector限制无匹配节点
- ❌ 污点容忍度不匹配
- ❌ 亲和性规则无法满足

**存储问题**
- ❌ PVC无法绑定到PV
- ❌ 存储类不可用

### 🌐 Kubernetes网络模型对比

| 网络插件 | 实现方式 | 性能 | 网络策略 | 适用场景 |
|---------|---------|------|---------|---------|
| **Flannel** | VXLAN/UDP Overlay | 一般 | ❌ | 小规模集群 |
| **Calico** | BGP纯三层 | 优秀 | ✅ | 大规模生产 |
| **Weave** | Overlay + 加密 | 一般 | ✅ | 安全要求高 |
| **Cilium** | eBPF | 极佳 | ✅ | 现代化集群 |

### 🗄️ ELK集群备份策略
```bash
# 1. 创建快照仓库
PUT /_snapshot/backup_repo
{
  "type": "fs",
  "settings": {
    "location": "/backup/elasticsearch"
  }
}

# 2. 创建快照
PUT /_snapshot/backup_repo/snapshot_1
{
  "indices": "*",
  "ignore_unavailable": true
}

# 3. 恢复快照
POST /_snapshot/backup_repo/snapshot_1/_restore
```

### 🐳 Docker vs Containerd

| 特性 | Docker | Containerd |
|------|--------|------------|
| **定位** | 完整容器平台 | 容器运行时 |
| **架构** | 分层架构 | 精简架构 |
| **性能** | 较重 | 轻量高效 |
| **K8s集成** | 通过dockershim | 原生CRI |
| **镜像管理** | Docker CLI | crictl |

### 🔌 网络插件实现原理
```yaml
# CNI插件工作流程
1. kubelet调用CNI插件
2. 创建网络命名空间
3. 创建veth pair
4. 配置IP地址和路由
5. 连接到网桥或隧道
```

### 📊 Prometheus自动发现
```yaml
# Kubernetes服务发现配置
scrape_configs:
  - job_name: 'kubernetes-pods'
    kubernetes_sd_configs:
    - role: pod
    relabel_configs:
    - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
      action: keep
      regex: true
    - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_port]
      action: replace
      target_label: __address__
      regex: (.+)
      replacement: ${1}
```

### 🔧 ES分片修改限制
- ✅ **副本分片**: 可动态修改
- ❌ **主分片**: 创建后不可修改
- 🔄 **解决方案**: Reindex API重建索引

### 🔴 Redis Cluster高可用
- **最少节点**: 6个（3主3从）
- **故障转移**: 自动主从切换
- **数据分片**: 16384个哈希槽
- **脑裂预防**: 需要奇数个主节点

### 🗑️ K8s节点删除流程
```bash
# 1. 标记不可调度
kubectl cordon <node-name>

# 2. 驱逐Pod
kubectl drain <node-name> --ignore-daemonsets --delete-emptydir-data

# 3. 删除节点
kubectl delete node <node-name>

# 4. 节点清理
kubeadm reset
```

---

## Go语言要点

### 🔄 Channel关闭检测
```go
// 方法1: ok语法
v, ok := <-ch
if !ok {
    // channel已关闭
}

// 方法2: range遍历（推荐）
for v := range ch {
    // 处理数据
}
```

### 🧵 Goroutine vs Thread
- **Goroutine**: 2KB初始栈，动态增长
- **Thread**: 8MB固定栈
- **调度**: GMP模型 vs 内核调度

### 🗑️ GC机制
- **三色标记**: 白色(待回收) → 灰色(标记中) → 黑色(存活)
- **写屏障**: 防止并发修改导致的错误回收
- **STW时间**: Go 1.8+ < 1ms

---

## 分布式系统

### 🎯 CAP理论
- **C (Consistency)**: 一致性
- **A (Availability)**: 可用性  
- **P (Partition tolerance)**: 分区容错性
- **定理**: 最多同时满足两个

### 🗳️ Raft算法核心
1. **Leader选举**: 心跳机制，任期概念
2. **日志复制**: 强一致性保证
3. **安全性**: Leader完整性原则

---

## 中间件技术

### 📨 Kafka一致性保证
```yaml
# 生产者配置
acks: all                    # 等待所有副本确认
retries: 2147483647         # 最大重试
enable.idempotence: true    # 幂等性
```

### 🔑 etcd核心特性
- **Watch机制**: 监听key变化
- **租约(Lease)**: TTL自动过期
- **事务**: 原子性操作

---

## 算法必备

### 🎯 高频题型
1. **数组**: 两数之和、三数之和、去重
2. **字符串**: 最长无重复子串、回文
3. **链表**: 反转、合并、环检测
4. **栈队列**: 有效括号、最小栈
5. **树**: 遍历、深度、对称性

### ⚡ 时间复杂度速记
- **O(1)**: 哈希表查找
- **O(log n)**: 二分搜索、堆操作
- **O(n)**: 线性扫描、BFS/DFS
- **O(n log n)**: 快排、归并排序
- **O(n²)**: 冒泡排序、暴力解法

---

## 🎪 考试答题模板

### 技术问题5步法
1. **理解确认** - "我理解的问题是..."
2. **思路分析** - "我的解决思路是..."
3. **代码实现** - "让我来实现一下..."
4. **测试验证** - "考虑边界情况..."
5. **优化改进** - "可以进一步优化..."

### 项目经验STAR法
- **S(Situation)**: 项目背景
- **T(Task)**: 具体任务
- **A(Action)**: 技术方案
- **R(Result)**: 最终效果

---

## 🔥 加分回答

### 展示深度思考
- "这个问题还可以从...角度考虑"
- "在生产环境中，我们还需要考虑..."
- "这种方案的trade-off是..."

### 体现实战经验
- "在我之前的项目中..."
- "我们遇到过类似问题，解决方案是..."
- "根据我的经验，最佳实践是..."

---

**💡 记住**: 保持自信，展示学习能力，诚实面对不会的问题！ 