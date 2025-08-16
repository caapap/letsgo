# 📊 监控告警面试题

> 适合：SRE工程师、运维开发、监控工程师
> 难度：⭐⭐⭐⭐ (中级-高级)

## 📋 监控系统设计

### 1. Prometheus架构

```yaml
# 核心组件
- Prometheus Server：时序数据库和查询引擎
- AlertManager：告警路由和通知
- Grafana：数据可视化和仪表板
- Pushgateway：推送指标接收器

# 高可用设计
- Prometheus集群化部署
- 数据分片和联邦集群
- 远程存储：Thanos/Cortex
- 负载均衡和故障转移
```

### 2. 监控指标体系

#### 黄金信号 (Golden Signals)
- **延迟 (Latency)**: 响应时间、P95/P99
- **流量 (Traffic)**: QPS、并发连接数
- **错误率 (Error Rate)**: 4xx/5xx错误
- **饱和度 (Saturation)**: 资源使用率

#### 监控层次
```
业务监控：用户活跃度、业务转化率
应用监控：应用性能、业务指标
基础设施：CPU、内存、磁盘、网络
```

### 3. 告警策略设计

#### 告警分类
- **P0**: 严重故障，立即响应
- **P1**: 重要问题，1小时内响应
- **P2**: 一般问题，4小时内响应
- **P3**: 轻微问题，24小时内响应

#### 告警规则优化
```yaml
# 避免告警风暴
groups:
- name: web-alerts
  rules:
  - alert: HighErrorRate
    expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.1
    for: 5m
    labels:
      severity: critical
    annotations:
      summary: "High error rate detected"
```

## 📋 可观测性

### 4. 三大支柱

#### Metrics（指标）
```prometheus
# 计数器
http_requests_total

# 直方图
http_request_duration_seconds

# 仪表盘
cpu_usage_percent
```

#### Logs（日志）
```json
{
  "timestamp": "2024-01-01T12:00:00Z",
  "level": "ERROR",
  "message": "Database connection failed",
  "trace_id": "abc123",
  "user_id": "user456"
}
```

#### Traces（追踪）
```yaml
# OpenTelemetry配置
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
exporters:
  jaeger:
    endpoint: jaeger:14250
```

### 5. 混沌工程

#### 故障注入类型
- **网络故障**: 延迟、丢包、分区
- **服务故障**: 服务不可用、响应缓慢
- **基础设施故障**: 节点宕机、磁盘满
- **依赖故障**: 数据库、缓存、消息队列

#### 工具选择
```bash
# Chaos Monkey
chaos-monkey --target-service web-app --failure-type latency

# Litmus
kubectl apply -f pod-delete-chaos.yaml

# Gremlin
gremlin attack-create shutdown --length 60 --target-type host
```

## 📋 监控实践

### 6. Prometheus vs OpenFalcon

| 特性 | Prometheus | OpenFalcon |
|------|------------|------------|
| 架构 | 拉取模式 | 推送模式 |
| 存储 | 自研时序数据库 | RRD格式 |
| 查询 | PromQL | 简单查询接口 |
| 生态 | CNCF项目 | 小米开源 |

### 7. 监控指标量级

```
系统指标：1000+ 指标/节点
应用指标：500+ 指标/服务
业务指标：200+ 指标/业务线
总计：10万+ 指标/集群
```

### 8. 日志架构

```yaml
# ELK Stack
Elasticsearch: 存储和搜索
Logstash: 数据处理
Kibana: 可视化

# 日志量级
日增量: 100GB - 1TB/天
总存储: 10TB - 100TB
索引数量: 100+ 索引
分片数量: 1000+ 分片
```

## 📋 告警管理

### 9. SLO指标设计

```yaml
# SLI（Service Level Indicator）
- 可用性：服务可用时间/总时间
- 延迟：响应时间分布（P50、P95、P99）
- 吞吐量：QPS、并发用户数
- 错误率：4xx、5xx错误比例

# SLO（Service Level Objective）
- 可用性目标：99.9%、99.99%
- 延迟目标：P95 < 200ms、P99 < 500ms
- 错误率目标：错误率 < 0.1%
```

### 10. 监控系统故障排查

```bash
# 常见故障
- Prometheus OOM：内存不足
- 数据丢失：磁盘空间不足
- 告警延迟：网络或配置问题
- 查询超时：数据量过大

# 排查方法
- 检查系统资源使用情况
- 查看日志和错误信息
- 验证配置和网络连接
- 测试数据采集和查询
```

## 🔗 相关资源

- [Prometheus官方文档](https://prometheus.io/docs/)
- [Grafana文档](https://grafana.com/docs/)
- [OpenTelemetry](https://opentelemetry.io/)

## 📝 复习要点

1. **掌握Prometheus架构和部署**
2. **理解监控指标体系设计**
3. **学会告警策略配置**
4. **熟悉可观测性三大支柱**
5. **实践混沌工程和故障排查**
