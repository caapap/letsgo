# 第一轮面试复习材料

## 📚 Kubernetes核心知识

### 1. CRD (Custom Resource Definition) 基础

#### 什么是CRD？
CRD允许用户在Kubernetes中定义自己的资源类型，扩展Kubernetes API。

#### CRD示例
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: appservices.app.example.com
spec:
  group: app.example.com
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              replicas:
                type: integer
              image:
                type: string
  scope: Namespaced
  names:
    plural: appservices
    singular: appservice
    kind: AppService
```

### 2. Operator开发框架

#### Kubebuilder vs Operator-SDK对比
| 特性 | Kubebuilder | Operator-SDK |
|------|------------|--------------|
| 语言支持 | Go | Go, Ansible, Helm |
| 学习曲线 | 中等 | 较低 |
| 社区支持 | Kubernetes官方 | RedHat/CoreOS |
| 脚手架 | 完善 | 更丰富 |

#### Controller核心组件
```go
// Reconcile核心逻辑
func (r *AppServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    // 1. 获取CR对象
    appService := &appv1.AppService{}
    if err := r.Get(ctx, req.NamespacedName, appService); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }
    
    // 2. 处理删除
    if !appService.DeletionTimestamp.IsZero() {
        return r.handleDeletion(ctx, appService)
    }
    
    // 3. 确保Deployment存在
    deployment := r.constructDeployment(appService)
    if err := r.createOrUpdate(ctx, deployment); err != nil {
        return ctrl.Result{}, err
    }
    
    // 4. 更新状态
    appService.Status.Ready = r.isDeploymentReady(deployment)
    if err := r.Status().Update(ctx, appService); err != nil {
        return ctrl.Result{}, err
    }
    
    return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}
```

### 3. Kubernetes网络知识

#### 网络模型
- **Cluster IP**: 集群内部访问
- **NodePort**: 节点端口暴露（30000-32767）
- **LoadBalancer**: 云厂商负载均衡器
- **Ingress**: 7层路由

#### 网络策略示例
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: backend-policy
spec:
  podSelector:
    matchLabels:
      app: backend
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: frontend
    ports:
    - protocol: TCP
      port: 8080
```

## 💻 Go语言重点

### 1. 并发模式

#### Worker Pool模式
```go
package main

import (
    "context"
    "sync"
)

type WorkerPool struct {
    workerCount int
    jobs        chan Job
    results     chan Result
    wg          sync.WaitGroup
}

type Job struct {
    ID   int
    Data interface{}
}

type Result struct {
    JobID int
    Data  interface{}
    Err   error
}

func NewWorkerPool(workerCount int) *WorkerPool {
    return &WorkerPool{
        workerCount: workerCount,
        jobs:        make(chan Job, workerCount*2),
        results:     make(chan Result, workerCount*2),
    }
}

func (p *WorkerPool) Start(ctx context.Context) {
    for i := 0; i < p.workerCount; i++ {
        p.wg.Add(1)
        go p.worker(ctx, i)
    }
}

func (p *WorkerPool) worker(ctx context.Context, id int) {
    defer p.wg.Done()
    
    for {
        select {
        case job, ok := <-p.jobs:
            if !ok {
                return
            }
            // 处理任务
            result := p.process(job)
            p.results <- result
            
        case <-ctx.Done():
            return
        }
    }
}

func (p *WorkerPool) process(job Job) Result {
    // 实际的业务逻辑
    return Result{JobID: job.ID}
}
```

### 2. 错误处理最佳实践

```go
// 自定义错误类型
type OpError struct {
    Op   string
    Kind string
    Err  error
}

func (e *OpError) Error() string {
    return fmt.Sprintf("operation %s failed on %s: %v", e.Op, e.Kind, e.Err)
}

func (e *OpError) Unwrap() error {
    return e.Err
}

// 错误包装
func doSomething() error {
    if err := callAPI(); err != nil {
        return &OpError{
            Op:   "doSomething",
            Kind: "API",
            Err:  err,
        }
    }
    return nil
}
```

## 🖥️ GPU资源管理基础

### NVIDIA GPU在Kubernetes中的使用

#### 1. Device Plugin部署
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: nvidia-device-plugin-daemonset
  namespace: kube-system
spec:
  selector:
    matchLabels:
      name: nvidia-device-plugin-ds
  template:
    spec:
      containers:
      - image: nvidia/k8s-device-plugin:v0.12.0
        name: nvidia-device-plugin-ctr
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop: ["ALL"]
        volumeMounts:
        - name: device-plugin
          mountPath: /var/lib/kubelet/device-plugins
```

#### 2. Pod请求GPU资源
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: gpu-pod
spec:
  containers:
  - name: cuda-container
    image: nvidia/cuda:11.0-base
    resources:
      limits:
        nvidia.com/gpu: 2 # 请求2个GPU
```

### GPU调度策略
- **独占模式**: 一个GPU只能被一个Pod使用
- **共享模式**: 通过MIG(Multi-Instance GPU)技术共享
- **时分复用**: 通过调度器实现GPU时间片共享

## 🔧 监控告警实践

### Prometheus配置示例
```yaml
# prometheus-config.yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'kubernetes-pods'
    kubernetes_sd_configs:
    - role: pod
    relabel_configs:
    - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
      action: keep
      regex: true
    - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
      action: replace
      target_label: __metrics_path__
      regex: (.+)
```

### 自定义告警规则
```yaml
groups:
- name: example
  rules:
  - alert: HighMemoryUsage
    expr: container_memory_usage_bytes{pod!=""} / container_spec_memory_limit_bytes > 0.9
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "Pod {{ $labels.pod }} memory usage is above 90%"
      description: "Pod {{ $labels.pod }} in namespace {{ $labels.namespace }} memory usage is {{ $value | humanizePercentage }}"
```

## 📝 面试常见问题速查

### Kubernetes相关
1. **Pod生命周期**: Pending → Running → Succeeded/Failed
2. **Service类型**: ClusterIP, NodePort, LoadBalancer, ExternalName
3. **存储类型**: emptyDir, hostPath, PV/PVC, ConfigMap, Secret
4. **调度策略**: nodeSelector, nodeAffinity, podAffinity, taints/tolerations

### Go语言相关
1. **并发原语**: goroutine, channel, select, sync包
2. **内存管理**: 栈分配vs堆分配, 逃逸分析
3. **GC优化**: GOGC设置, sync.Pool使用
4. **性能分析**: pprof, trace, benchmark

### 网络相关
1. **OSI七层模型**: 物理层→数据链路层→网络层→传输层→会话层→表示层→应用层
2. **TCP三次握手**: SYN → SYN-ACK → ACK
3. **Linux网络命令**: tcpdump, netstat, ss, ip, iptables

## 🎯 实战项目准备建议

### 项目案例模板
```
项目名称：XXX自动化运维平台
项目背景：
- 管理规模：500+节点的Kubernetes集群
- 日均处理：10万+容器调度请求
- 服务对象：50+研发团队

技术方案：
- 架构：微服务架构，Go语言开发
- 核心组件：自定义Operator、调度器、监控告警
- 技术栈：Go + Kubernetes + Prometheus + Grafana

项目成果：
- 部署效率提升300%
- 故障恢复时间从小时级降到分钟级
- 资源利用率提升40%

个人贡献：
- 负责Operator开发
- 设计自动扩缩容策略
- 优化调度算法
```

## 下一步学习计划
1. 深入学习Kubernetes源码
2. 实践CRD/Operator开发项目
3. 了解GPU虚拟化技术
4. 学习服务网格(Istio/Linkerd)
5. 掌握GitOps实践(ArgoCD/Flux) 