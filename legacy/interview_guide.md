# 算力服务平台运维开发岗位面试指导

## 1. Go语言基础与高级特性

### 1.1 基础问题
**切片(slice)与数组(array)的区别**
- **数组**: 固定长度，值类型，内存连续
- **切片**: 动态数组，引用类型，底层是数组的引用
```go
// 数组
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// 切片
var slice []int = []int{1, 2, 3, 4, 5}
slice = append(slice, 6) // 可以动态扩容
```

**Channel关闭判断**
```go
// 方法1: 使用ok语法
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    close(ch)
}()

for {
    v, ok := <-ch
    if !ok {
        fmt.Println("Channel已关闭")
        break
    }
    fmt.Println("接收到:", v)
}

// 方法2: 使用range (推荐)
for v := range ch {
    fmt.Println("接收到:", v)
}
```

### 1.2 GDB调试器相关
```bash
# 编译时保留调试信息
go build -gcflags "-N -l" main.go

# GDB常用命令
gdb ./main
(gdb) b main.main    # 设置断点
(gdb) r              # 运行程序  
(gdb) n              # 下一行
(gdb) s              # 步入函数
(gdb) p variable     # 打印变量值
(gdb) bt             # 查看调用栈
```

### 1.3 Go并发编程
**Goroutine vs Thread**
- Goroutine: 轻量级线程，初始栈2KB，动态增长
- 系统线程: 重量级，初始栈8MB

**并发模式**
```go
// Worker Pool模式
func workerPool(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        results <- job * 2
    }
}

// 使用示例
jobs := make(chan int, 100)
results := make(chan int, 100)
var wg sync.WaitGroup

// 启动3个worker
for i := 0; i < 3; i++ {
    wg.Add(1)
    go workerPool(jobs, results, &wg)
}

// 发送任务
for i := 1; i <= 10; i++ {
    jobs <- i
}
close(jobs)

wg.Wait()
close(results)
```

## 2. 分布式系统与中间件

### 2.1 Kafka消息一致性保证
**生产者一致性**
```yaml
# 关键配置
acks: all                    # 等待所有副本确认
retries: 2147483647         # 最大重试次数
enable.idempotence: true    # 开启幂等性
max.in.flight.requests.per.connection: 5
```

**消费者一致性**
```go
// 手动提交offset保证一致性
consumer := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servers": "localhost:9092",
    "group.id":          "my-group",
    "enable.auto.commit": false,  // 关闭自动提交
})

for {
    msg, err := consumer.ReadMessage(-1)
    if err == nil {
        // 处理消息
        processMessage(msg)
        
        // 手动提交offset
        consumer.CommitMessage(msg)
    }
}
```

### 2.2 etcd数据同步
**Raft一致性算法**
- Leader选举: 心跳机制，任期概念
- 日志复制: 强一致性保证
- 安全性: Leader完整性原则

```go
// etcd客户端使用
client, err := clientv3.New(clientv3.Config{
    Endpoints:   []string{"localhost:2379"},
    DialTimeout: 5 * time.Second,
})

// 监听key变化
watchCh := client.Watch(context.Background(), "/config/")
for resp := range watchCh {
    for _, event := range resp.Events {
        fmt.Printf("事件类型: %s, Key: %s, Value: %s\n", 
            event.Type, event.Kv.Key, event.Kv.Value)
    }
}
```

## 3. Kubernetes集群建设

### 3.1 核心组件
**控制平面组件**
- **kube-apiserver**: API网关，所有操作入口
- **kube-controller-manager**: 控制器管理器
- **kube-scheduler**: 调度器
- **etcd**: 分布式键值存储

**节点组件**
- **kubelet**: 节点代理
- **kube-proxy**: 网络代理
- **容器运行时**: Docker/containerd

### 3.2 网络模型与插件
**K8s网络模型要求**
1. Pod内容器共享网络
2. Pod间可直接通信（无NAT）
3. Node与Pod可直接通信
4. Pod看到的IP就是其他Pod看到的IP

**常用CNI插件对比**
```yaml
# Flannel - 简单overlay网络
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-flannel-cfg
data:
  net-conf.json: |
    {
      "Network": "10.244.0.0/16",
      "Backend": {
        "Type": "vxlan"
      }
    }

# Calico - 性能更好，支持网络策略
apiVersion: projectcalico.org/v3
kind: NetworkPolicy
metadata:
  name: deny-all
spec:
  podSelector: {}
  policyTypes:
  - Ingress
  - Egress
```

### 3.3 大规模集群建设
**集群规模考虑**
- 单集群最大5000节点
- 每节点最大110个Pod
- 集群总Pod数不超过150,000

**高可用部署**
```yaml
# 多Master部署
apiVersion: kubeadm.k8s.io/v1beta3
kind: ClusterConfiguration
kubernetesVersion: v1.28.0
controlPlaneEndpoint: "loadbalancer.example.com:6443"
etcd:
  external:
    endpoints:
    - "https://etcd1.example.com:2379"
    - "https://etcd2.example.com:2379"
    - "https://etcd3.example.com:2379"
```

## 4. 算法题详解

### 4.1 数组去重（原地修改）
```go
// 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}

// 测试
func main() {
    nums := []int{1, 1, 2, 2, 3, 4, 4}
    newLen := removeDuplicates(nums)
    fmt.Printf("新长度: %d, 数组: %v\n", newLen, nums[:newLen])
    // 输出: 新长度: 4, 数组: [1 2 3 4]
}
```

### 4.2 多数元素（出现次数>n/2）
```go
// Boyer-Moore投票算法
func majorityElement(nums []int) int {
    candidate := nums[0]
    count := 1
    
    // 投票阶段
    for i := 1; i < len(nums); i++ {
        if count == 0 {
            candidate = nums[i]
            count = 1
        } else if nums[i] == candidate {
            count++
        } else {
            count--
        }
    }
    
    return candidate
}

// 时间复杂度: O(n), 空间复杂度: O(1)
```

### 4.3 最长无重复字符子串
```go
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    charMap := make(map[byte]int)
    left, maxLen := 0, 0
    
    for right := 0; right < len(s); right++ {
        // 如果字符已存在且在当前窗口内
        if pos, exists := charMap[s[right]]; exists && pos >= left {
            left = pos + 1
        }
        
        charMap[s[right]] = right
        maxLen = max(maxLen, right-left+1)
    }
    
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 示例: "abcabcbb" -> 3 ("abc")
```

### 4.4 两两交换链表节点
```go
type ListNode struct {
    Val  int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    // 创建虚拟头节点
    dummy := &ListNode{Next: head}
    prev := dummy
    
    for prev.Next != nil && prev.Next.Next != nil {
        // 保存要交换的两个节点
        first := prev.Next
        second := prev.Next.Next
        
        // 执行交换
        prev.Next = second
        first.Next = second.Next
        second.Next = first
        
        // 移动prev指针
        prev = first
    }
    
    return dummy.Next
}

// 1->2->3->4 变成 2->1->4->3
```

### 4.5 LeetCode高频题目清单
**数组/字符串**
1. 两数之和 (Two Sum)
2. 三数之和 (3Sum)
3. 最长公共前缀 (Longest Common Prefix)
4. 合并两个有序数组 (Merge Sorted Array)

**链表**
5. 反转链表 (Reverse Linked List)
6. 合并两个有序链表 (Merge Two Sorted Lists)
7. 环形链表 (Linked List Cycle)

**栈/队列**
8. 有效的括号 (Valid Parentheses)
9. 最小栈 (Min Stack)

**树**
10. 二叉树的最大深度 (Maximum Depth of Binary Tree)
11. 对称二叉树 (Symmetric Tree)

## 5. 面试软技能

### 5.1 项目沟通技巧
**STAR法则描述项目**
- **Situation**: 项目背景
- **Task**: 具体任务
- **Action**: 采取的行动
- **Result**: 取得的结果

### 5.2 技术问题回答策略
1. **先说思路**: 整体架构设计思路
2. **细化方案**: 具体技术选型和实现
3. **考虑边界**: 异常情况处理
4. **性能优化**: 可扩展性和性能考虑

### 5.3 提问环节
**了解岗位职责**
- "这个岗位的主要技术栈是什么？"
- "团队规模和协作方式是怎样的？"
- "算力平台的核心业务场景有哪些？"

**技术发展方向**
- "公司在云原生技术方面的发展规划？"
- "团队对新技术的学习和应用机制？"

## 6. 算力服务平台特定知识点

### 6.1 GPU/NPU资源管理
```yaml
# GPU资源调度
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: gpu-container
    image: nvidia/cuda:11.0-base
    resources:
      limits:
        nvidia.com/gpu: 2  # 申请2个GPU
```

### 6.2 大规模任务调度
**调度策略**
- 资源感知调度
- 优先级队列
- 抢占式调度
- 亲和性调度

### 6.3 监控告警体系
```yaml
# Prometheus监控配置
global:
  scrape_interval: 15s
  evaluation_interval: 15s

rule_files:
  - "gpu_rules.yml"
  - "cluster_rules.yml"

scrape_configs:
  - job_name: 'kubernetes-nodes'
    kubernetes_sd_configs:
    - role: node
```

这份指导涵盖了算力服务平台运维开发岗位的核心技术点，重点关注Go语言、Kubernetes、分布式系统和算法。建议按模块重点复习，特别是Go并发、K8s网络和分布式一致性等高频考点。 