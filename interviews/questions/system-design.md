# 🏗️ 系统设计面试题

> 适合：高级工程师、架构师、SRE工程师
> 难度：⭐⭐⭐⭐⭐ (高级)

## 📋 大规模系统设计

### 1. 监控系统架构

#### 百万级指标监控系统
```yaml
架构分层：
数据采集层: Agent、Exporter、Push Gateway
数据传输层: 消息队列、负载均衡
数据存储层: 时序数据库集群、分片策略
查询计算层: 查询引擎、缓存层
展示告警层: Dashboard、告警系统

高可用设计:
- 多机房部署：异地多活
- 服务冗余：每个组件至少3个实例
- 故障转移：自动故障检测和切换
- 数据备份：多副本、定期备份

性能优化:
- 数据分片：按时间、标签分片
- 读写分离：写入优化、查询优化
- 缓存策略：多级缓存、智能预取
- 异步处理：非阻塞操作、批量处理
```

### 2. 微服务可观测性

#### 分布式追踪
```yaml
# OpenTelemetry配置
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
processors:
  batch:
    timeout: 1s
    send_batch_size: 1024
exporters:
  jaeger:
    endpoint: jaeger:14250
service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [jaeger]
```

#### 三大支柱整合
- **Metrics**: Prometheus、StatsD
- **Logs**: ELK Stack、Loki
- **Traces**: Jaeger、Zipkin

### 3. 容量规划

#### 系统容量计算
```
业务指标分析:
- 用户增长预测：DAU、MAU增长率
- 业务峰值：节假日、促销活动
- 功能复杂度：新功能对资源的影响

技术指标计算:
- 峰值QPS = 日活用户 × 平均请求次数 ÷ 86400 × 峰值系数
- 存储容量 = 用户数 × 平均数据量 × 增长系数
- 带宽需求 = 请求数 × 平均响应大小 × 并发系数
- 计算资源 = CPU密集型 + 内存密集型 + I/O密集型
```

## 📋 高可用架构

### 4. 多活架构设计

#### 机房数量选择
- **双活**: 2个机房，成本低但风险高
- **三活**: 3个机房，推荐方案，平衡成本风险
- **四活**: 4个机房，高可用但成本高
- **五活**: 5个机房，极高可用，成本最高

#### 设计考虑因素
```yaml
业务连续性要求: RTO、RPO
成本预算: 建设成本、运维成本
技术复杂度: 数据同步、网络延迟
运维能力: 团队技能、工具支持
```

### 5. 故障隔离设计

#### 隔离策略
```go
// 熔断器模式
type CircuitBreaker struct {
    state       State
    failureThreshold int64
    timeout     time.Duration
    lastFailure time.Time
    mutex       sync.RWMutex
}

func (cb *CircuitBreaker) Execute(command func() error) error {
    if cb.ReadyToTrip() {
        return ErrCircuitBreakerOpen
    }
    
    err := command()
    if err != nil {
        cb.recordFailure()
    }
    return err
}
```

#### 降级策略
- **功能降级**: 非核心功能关闭
- **服务降级**: 备用服务、简化逻辑
- **数据降级**: 缓存数据、历史数据
- **用户体验降级**: 简化界面、减少交互

## 📋 负载均衡

### 6. Nginx负载均衡

#### 调度算法
```nginx
# 轮询（默认）
upstream backend {
    server 192.168.1.10:8080;
    server 192.168.1.11:8080;
}

# 加权轮询
upstream backend {
    server 192.168.1.10:8080 weight=3;
    server 192.168.1.11:8080 weight=1;
}

# IP哈希
upstream backend {
    ip_hash;
    server 192.168.1.10:8080;
    server 192.168.1.11:8080;
}

# 最少连接
upstream backend {
    least_conn;
    server 192.168.1.10:8080;
    server 192.168.1.11:8080;
}
```

### 7. LVS负载均衡

#### 工作模式对比
| 模式 | 原理 | 优势 | 劣势 |
|------|------|------|------|
| NAT | 修改IP地址 | 配置简单 | 性能瓶颈 |
| DR | 修改MAC地址 | 性能高 | 配置复杂 |
| TUN | IP隧道 | 可跨网段 | 最复杂 |

## 📋 数据库设计

### 8. MySQL优化

#### 索引设计
```sql
-- B+树索引优势
- 范围查询效率高：叶子节点链表结构
- 稳定的查询性能：所有查询都到叶子节点
- 更好的缓存局部性：内部节点更小

-- 索引类型
CREATE INDEX idx_name ON users(name);              -- 普通索引
CREATE UNIQUE INDEX idx_email ON users(email);     -- 唯一索引
CREATE INDEX idx_name_email ON users(name, email); -- 复合索引
```

#### 事务隔离级别
- **READ UNCOMMITTED**: 读未提交
- **READ COMMITTED**: 读已提交
- **REPEATABLE READ**: 可重复读（MySQL默认）
- **SERIALIZABLE**: 串行化

## 📋 容器化架构

### 9. Docker优化

#### 镜像优化
```dockerfile
# 多阶段构建
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /usr/local/bin/
CMD ["main"]

# 层优化原则
RUN apt-get update && \
    apt-get install -y package && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
```

### 10. SSL/TLS安全

#### 握手过程
```
1. 客户端Hello: 发送支持的TLS版本、随机数、加密套件
2. 服务器Hello: 选择TLS版本、发送随机数、证书
3. 客户端验证: 验证证书有效性、域名匹配
4. 密钥交换: 生成预主密钥，用公钥加密发送
5. 生成会话密钥: 双方生成对称加密密钥
6. 握手完成: 开始对称加密通信

为什么先非对称后对称？
- 非对称加密：安全但慢，用于密钥交换
- 对称加密：快速但需要共享密钥，用于数据传输
```

## 🔗 相关资源

- [系统设计面试指南](https://github.com/donnemartin/system-design-primer)
- [微服务架构设计](https://microservices.io/)
- [高可用架构](https://aws.amazon.com/architecture/reliability/)

## 📝 复习要点

1. **掌握大规模系统架构设计原则**
2. **理解微服务可观测性设计方法**
3. **学会容量规划和性能评估**
4. **熟悉高可用和故障隔离策略**
5. **实践负载均衡和数据库优化**
