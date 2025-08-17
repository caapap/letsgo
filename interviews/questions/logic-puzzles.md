# 🧩 逻辑思维与算法题

> 适合：所有技术岗位、锻炼逻辑思维
> 难度：⭐⭐⭐⭐ (中级-高级)
> 来源：[百度、字节跳动等大厂面试真题](https://blog.csdn.net/CCIEHL/article/details/104151990)

## 📋 逻辑推理题

### 1. 找出最轻的球

**问题**: 16个球，只有一个更轻，一个天平，怎么三次找出最轻的球？

**解答思路**:
```
第一次称重：
- 将16个球分成三组：6、6、4
- 称重两个6球组
- 如果平衡，轻球在4球组中
- 如果不平衡，轻球在较轻的6球组中

第二次称重：
- 情况1：如果轻球在4球组
  - 分成2、2称重
  - 找出较轻的2球组
- 情况2：如果轻球在6球组
  - 分成2、2、2
  - 称重两个2球组
  - 确定轻球所在的2球组

第三次称重：
- 将剩余的2个球称重
- 找出较轻的球

优化方案（更通用）：
- 使用三分法
- 每次将球分成尽可能相等的三组
- 通过称重排除2/3的可能性
```

**关键点**:
- 利用天平的三种状态：左重、右重、平衡
- 每次称重最大化信息获取
- 分组策略的优化

---

### 2. 烧绳计时问题

**问题**: 烧一根不均匀的绳子，从头烧到尾是要1个小时。现在有若干条材质相同的绳子，问如何用烧绳的方法来计时一个小时15分钟？

**解答思路**:
```
基础知识：
- 一根绳子烧完需要1小时
- 同时点燃两端，烧完需要30分钟
- 绳子不均匀，但总时间固定

解决方案：
1. 准备两根绳子A和B
2. T=0时刻：
   - 绳子A：同时点燃两端
   - 绳子B：只点燃一端
3. T=30分钟（A烧完）：
   - 绳子B：此时还剩30分钟
   - 立即点燃B的另一端
4. T=45分钟（B烧完）：
   - 已经过了30+15=45分钟
5. 再点燃一根新绳子C：
   - 只点燃一端
6. T=1小时45分钟（C烧完）：
   - 总计时：45分钟 + 60分钟 = 1小时15分钟

更简单的方案：
1. 同时开始：
   - 绳子A：点燃一端（60分钟）
   - 绳子B：点燃两端（30分钟）
2. B烧完后（30分钟）：
   - 绳子C：点燃两端（15分钟）
   - A继续烧
3. C烧完后（45分钟）：
   - A继续烧
4. A烧完（75分钟）：
   - 总时间：1小时15分钟
```

**扩展思考**:
- 如何计时45分钟？
- 如何计时任意时间？
- 最少需要几根绳子？

---

## 📋 算法编程题

### 3. B+树实现

**问题**: 选一个喜欢的语言，写B+树

**Go语言实现**:
```go
package main

type BPlusNode struct {
    isLeaf   bool
    keys     []int
    children []*BPlusNode
    next     *BPlusNode  // 叶子节点链表
    parent   *BPlusNode
}

type BPlusTree struct {
    root  *BPlusNode
    order int  // B+树的阶数
}

func NewBPlusTree(order int) *BPlusTree {
    return &BPlusTree{
        root:  &BPlusNode{isLeaf: true},
        order: order,
    }
}

// 插入操作
func (t *BPlusTree) Insert(key int) {
    root := t.root
    
    // 如果根节点满了，需要分裂
    if len(root.keys) >= t.order-1 {
        newRoot := &BPlusNode{isLeaf: false}
        newRoot.children = append(newRoot.children, root)
        t.splitChild(newRoot, 0)
        t.root = newRoot
        root = newRoot
    }
    
    t.insertNonFull(root, key)
}

// 在非满节点插入
func (t *BPlusTree) insertNonFull(node *BPlusNode, key int) {
    i := len(node.keys) - 1
    
    if node.isLeaf {
        // 叶子节点直接插入
        node.keys = append(node.keys, 0)
        for i >= 0 && node.keys[i] > key {
            node.keys[i+1] = node.keys[i]
            i--
        }
        node.keys[i+1] = key
    } else {
        // 非叶子节点，找到合适的子节点
        for i >= 0 && node.keys[i] > key {
            i--
        }
        i++
        
        if len(node.children[i].keys) >= t.order-1 {
            t.splitChild(node, i)
            if key > node.keys[i] {
                i++
            }
        }
        t.insertNonFull(node.children[i], key)
    }
}

// 分裂子节点
func (t *BPlusTree) splitChild(parent *BPlusNode, index int) {
    fullNode := parent.children[index]
    newNode := &BPlusNode{isLeaf: fullNode.isLeaf}
    
    mid := t.order / 2
    
    // 分配键值
    newNode.keys = append(newNode.keys, fullNode.keys[mid:]...)
    fullNode.keys = fullNode.keys[:mid]
    
    // 如果不是叶子节点，分配子节点
    if !fullNode.isLeaf {
        newNode.children = append(newNode.children, fullNode.children[mid:]...)
        fullNode.children = fullNode.children[:mid]
    }
    
    // 更新父节点
    parent.keys = append(parent.keys[:index], 
        append([]int{fullNode.keys[mid-1]}, parent.keys[index:]...)...)
    parent.children = append(parent.children[:index+1], 
        append([]*BPlusNode{newNode}, parent.children[index+1:]...)...)
    
    // 如果是叶子节点，维护链表
    if fullNode.isLeaf {
        newNode.next = fullNode.next
        fullNode.next = newNode
    }
}

// 查找操作
func (t *BPlusTree) Search(key int) bool {
    return t.searchNode(t.root, key)
}

func (t *BPlusTree) searchNode(node *BPlusNode, key int) bool {
    i := 0
    for i < len(node.keys) && key > node.keys[i] {
        i++
    }
    
    if i < len(node.keys) && key == node.keys[i] {
        if node.isLeaf {
            return true
        }
    }
    
    if node.isLeaf {
        return false
    }
    
    return t.searchNode(node.children[i], key)
}
```

**Python实现**:
```python
class BPlusNode:
    def __init__(self, is_leaf=True):
        self.keys = []
        self.children = []
        self.is_leaf = is_leaf
        self.next = None  # 叶子节点链表
        
class BPlusTree:
    def __init__(self, order=3):
        self.root = BPlusNode()
        self.order = order
    
    def insert(self, key):
        root = self.root
        if len(root.keys) >= self.order - 1:
            new_root = BPlusNode(is_leaf=False)
            new_root.children.append(root)
            self._split_child(new_root, 0)
            self.root = new_root
            root = new_root
        self._insert_non_full(root, key)
    
    def search(self, key):
        return self._search_node(self.root, key)
```

---

### 4. 快速排序实现

**问题**: 写一个快排

**Go语言实现**:
```go
func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    pivot := arr[len(arr)/2]
    left := []int{}
    right := []int{}
    equal := []int{}
    
    for _, v := range arr {
        if v < pivot {
            left = append(left, v)
        } else if v > pivot {
            right = append(right, v)
        } else {
            equal = append(equal, v)
        }
    }
    
    left = quickSort(left)
    right = quickSort(right)
    
    return append(append(left, equal...), right...)
}

// 原地快排（更高效）
func quickSortInPlace(arr []int, low, high int) {
    if low < high {
        pi := partition(arr, low, high)
        quickSortInPlace(arr, low, pi-1)
        quickSortInPlace(arr, pi+1, high)
    }
}

func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}
```

---

### 5. 单链表实现

**问题**: Python实现一个单链表

**Python实现**:
```python
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class LinkedList:
    def __init__(self):
        self.head = None
        self.size = 0
    
    def add_first(self, val):
        """在链表头部添加节点"""
        new_node = ListNode(val)
        new_node.next = self.head
        self.head = new_node
        self.size += 1
    
    def add_last(self, val):
        """在链表尾部添加节点"""
        new_node = ListNode(val)
        if not self.head:
            self.head = new_node
        else:
            current = self.head
            while current.next:
                current = current.next
            current.next = new_node
        self.size += 1
    
    def remove_first(self):
        """删除头节点"""
        if not self.head:
            return None
        val = self.head.val
        self.head = self.head.next
        self.size -= 1
        return val
    
    def remove(self, val):
        """删除指定值的节点"""
        if not self.head:
            return False
        
        if self.head.val == val:
            self.head = self.head.next
            self.size -= 1
            return True
        
        current = self.head
        while current.next:
            if current.next.val == val:
                current.next = current.next.next
                self.size -= 1
                return True
            current = current.next
        return False
    
    def reverse(self):
        """反转链表"""
        prev = None
        current = self.head
        
        while current:
            next_node = current.next
            current.next = prev
            prev = current
            current = next_node
        
        self.head = prev
    
    def find_middle(self):
        """找到链表中间节点（快慢指针）"""
        if not self.head:
            return None
        
        slow = fast = self.head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        
        return slow.val
    
    def has_cycle(self):
        """检测链表是否有环"""
        if not self.head:
            return False
        
        slow = fast = self.head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return True
        
        return False
    
    def display(self):
        """打印链表"""
        result = []
        current = self.head
        while current:
            result.append(str(current.val))
            current = current.next
        return " -> ".join(result)
```

---

### 6. 装饰器实现重试功能

**问题**: Python装饰器实现重试功能代码实现

**Python实现**:
```python
import time
import functools
from typing import Callable, Any, Optional, Type, Tuple

def retry(
    max_attempts: int = 3,
    delay: float = 1.0,
    backoff: float = 2.0,
    exceptions: Tuple[Type[Exception], ...] = (Exception,),
    on_retry: Optional[Callable] = None
):
    """
    重试装饰器
    
    Args:
        max_attempts: 最大重试次数
        delay: 初始延迟时间（秒）
        backoff: 延迟时间的倍数（指数退避）
        exceptions: 需要捕获的异常类型
        on_retry: 重试时的回调函数
    """
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs) -> Any:
            attempt = 1
            current_delay = delay
            
            while attempt <= max_attempts:
                try:
                    return func(*args, **kwargs)
                except exceptions as e:
                    if attempt == max_attempts:
                        print(f"最终失败: {func.__name__} 在 {max_attempts} 次尝试后失败")
                        raise
                    
                    if on_retry:
                        on_retry(func.__name__, attempt, e)
                    else:
                        print(f"重试 {attempt}/{max_attempts}: {func.__name__} 失败: {e}")
                    
                    time.sleep(current_delay)
                    current_delay *= backoff
                    attempt += 1
            
            return None
        
        return wrapper
    return decorator

# 使用示例
@retry(max_attempts=3, delay=1, backoff=2)
def unstable_api_call():
    """模拟不稳定的API调用"""
    import random
    if random.random() < 0.7:  # 70%概率失败
        raise ConnectionError("API连接失败")
    return "成功"

# 更高级的重试装饰器
class RetryWithContext:
    """带上下文的重试装饰器"""
    
    def __init__(
        self,
        max_attempts: int = 3,
        delay: float = 1.0,
        backoff: float = 2.0,
        max_delay: float = 60.0,
        exceptions: Tuple[Type[Exception], ...] = (Exception,),
        should_retry: Optional[Callable[[Exception], bool]] = None
    ):
        self.max_attempts = max_attempts
        self.delay = delay
        self.backoff = backoff
        self.max_delay = max_delay
        self.exceptions = exceptions
        self.should_retry = should_retry
    
    def __call__(self, func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs) -> Any:
            attempt = 1
            current_delay = self.delay
            last_exception = None
            
            while attempt <= self.max_attempts:
                try:
                    result = func(*args, **kwargs)
                    if attempt > 1:
                        print(f"成功: {func.__name__} 在第 {attempt} 次尝试成功")
                    return result
                    
                except self.exceptions as e:
                    last_exception = e
                    
                    # 检查是否应该重试
                    if self.should_retry and not self.should_retry(e):
                        raise
                    
                    if attempt == self.max_attempts:
                        raise
                    
                    # 记录重试信息
                    print(f"重试 {attempt}/{self.max_attempts}: "
                          f"{func.__name__} 失败: {e.__class__.__name__}: {e}")
                    
                    # 等待并增加延迟
                    time.sleep(current_delay)
                    current_delay = min(current_delay * self.backoff, self.max_delay)
                    attempt += 1
            
            raise last_exception
        
        return wrapper

# 使用示例
def should_retry_on_rate_limit(e: Exception) -> bool:
    """判断是否应该重试"""
    if isinstance(e, ValueError) and "rate limit" in str(e).lower():
        return True
    return isinstance(e, (ConnectionError, TimeoutError))

@RetryWithContext(
    max_attempts=5,
    delay=1,
    backoff=2,
    max_delay=30,
    exceptions=(ConnectionError, TimeoutError, ValueError),
    should_retry=should_retry_on_rate_limit
)
def api_request(url: str):
    """带重试的API请求"""
    # 实际的API请求逻辑
    pass

# 异步重试装饰器
import asyncio

def async_retry(max_attempts: int = 3, delay: float = 1.0):
    """异步重试装饰器"""
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        async def wrapper(*args, **kwargs):
            for attempt in range(1, max_attempts + 1):
                try:
                    return await func(*args, **kwargs)
                except Exception as e:
                    if attempt == max_attempts:
                        raise
                    print(f"异步重试 {attempt}/{max_attempts}: {e}")
                    await asyncio.sleep(delay * attempt)
            return None
        return wrapper
    return decorator
```

---

## 📋 故障排查题

### 7. 用户无法访问问题

**问题**: 一个用户不能访问，解决思路

**排查流程**:
```yaml
1. 收集信息:
   - 用户位置（地域、网络环境）
   - 访问时间
   - 错误信息/现象
   - 影响范围（个别用户还是批量）

2. 网络层排查:
   - ping测试连通性
   - traceroute查看路由路径
   - nslookup/dig检查DNS解析
   - telnet测试端口连通性

3. 应用层排查:
   - 检查服务状态
   - 查看应用日志
   - 检查负载均衡器
   - 验证SSL证书

4. 用户侧排查:
   - 浏览器缓存/Cookie
   - 代理设置
   - 防火墙/安全软件
   - DNS设置

5. 具体命令:
```bash
# 网络连通性
ping -c 10 example.com
traceroute example.com
mtr example.com

# DNS检查
nslookup example.com
dig example.com @8.8.8.8

# 端口测试
telnet example.com 443
nc -zv example.com 443

# HTTP测试
curl -v https://example.com
curl -I https://example.com

# 证书检查
openssl s_client -connect example.com:443
```

---

### 8. 用户劫持检测

**问题**: 怎么证明一个用户被劫持了？

**检测方法**:
```yaml
1. DNS劫持检测:
   - 对比不同DNS服务器解析结果
   - 检查解析IP是否在预期范围
   - 监控DNS TTL异常

2. HTTP劫持检测:
   - 检查响应内容是否被篡改
   - 验证响应头是否正确
   - 检测是否有注入的JS/广告

3. 证书劫持检测:
   - 验证SSL证书链
   - 检查证书指纹
   - 对比证书颁发机构

4. 检测脚本:
```bash
#!/bin/bash
# DNS劫持检测
domain="example.com"
expected_ip="1.2.3.4"

# 测试多个DNS
for dns in 8.8.8.8 114.114.114.114 1.1.1.1; do
    result=$(dig @$dns $domain +short)
    if [ "$result" != "$expected_ip" ]; then
        echo "DNS劫持检测: $dns 返回 $result"
    fi
done

# HTTP劫持检测
response=$(curl -s https://example.com)
if echo "$response" | grep -q "injected_content"; then
    echo "检测到内容注入"
fi

# 证书检测
cert_fingerprint=$(echo | openssl s_client -connect example.com:443 2>/dev/null | \
    openssl x509 -fingerprint -noout)
expected_fingerprint="SHA1 Fingerprint=XX:XX:XX..."
if [ "$cert_fingerprint" != "$expected_fingerprint" ]; then
    echo "证书可能被替换"
fi
```

---

### 9. 系统监控工具对比

**问题**: iotop iostat区别

**详细对比**:
```yaml
iostat:
- 功能：显示CPU和磁盘I/O统计信息
- 级别：系统级别统计
- 输出：平均值和累计值
- 用途：整体I/O性能分析

iotop:
- 功能：显示进程级别的I/O使用情况
- 级别：进程级别实时监控
- 输出：每个进程的读写速率
- 用途：找出I/O密集型进程

使用示例:
```bash
# iostat - 系统级I/O统计
iostat -x 1 10  # 每秒刷新，共10次
iostat -d -p sda  # 显示sda磁盘详细信息

# iotop - 进程级I/O监控
iotop -o  # 只显示有I/O的进程
iotop -P  # 显示进程而非线程
iotop -a  # 显示累计I/O

# 其他相关工具
dstat  # 综合监控工具
atop   # 高级系统监控
htop   # 增强版top
vmstat # 虚拟内存统计
```

---

### 10. 搭建HTTP DNS

**问题**: 如何自己搭建一个httpdns

**实现方案**:
```go
// HTTP DNS服务器实现
package main

import (
    "encoding/json"
    "fmt"
    "net"
    "net/http"
    "sync"
    "time"
)

type DNSCache struct {
    sync.RWMutex
    records map[string]*DNSRecord
}

type DNSRecord struct {
    IPs       []string  `json:"ips"`
    TTL       int       `json:"ttl"`
    UpdatedAt time.Time `json:"updated_at"`
}

type HTTPDNSServer struct {
    cache     *DNSCache
    upstream  string
}

func NewHTTPDNSServer() *HTTPDNSServer {
    return &HTTPDNSServer{
        cache: &DNSCache{
            records: make(map[string]*DNSRecord),
        },
        upstream: "8.8.8.8:53",
    }
}

// HTTP接口处理
func (s *HTTPDNSServer) HandleResolve(w http.ResponseWriter, r *http.Request) {
    domain := r.URL.Query().Get("domain")
    if domain == "" {
        http.Error(w, "domain parameter required", http.StatusBadRequest)
        return
    }
    
    // 查询缓存
    record := s.getFromCache(domain)
    if record == nil {
        // 缓存未命中，进行DNS查询
        ips, err := s.resolve(domain)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        record = &DNSRecord{
            IPs:       ips,
            TTL:       300,
            UpdatedAt: time.Now(),
        }
        
        s.updateCache(domain, record)
    }
    
    // 返回JSON结果
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(record)
}

// DNS解析
func (s *HTTPDNSServer) resolve(domain string) ([]string, error) {
    ips, err := net.LookupHost(domain)
    if err != nil {
        return nil, err
    }
    return ips, nil
}

// 缓存管理
func (s *HTTPDNSServer) getFromCache(domain string) *DNSRecord {
    s.cache.RLock()
    defer s.cache.RUnlock()
    
    record, exists := s.cache.records[domain]
    if !exists {
        return nil
    }
    
    // 检查TTL
    if time.Since(record.UpdatedAt).Seconds() > float64(record.TTL) {
        return nil
    }
    
    return record
}

func (s *HTTPDNSServer) updateCache(domain string, record *DNSRecord) {
    s.cache.Lock()
    defer s.cache.Unlock()
    s.cache.records[domain] = record
}

// 启动服务器
func main() {
    server := NewHTTPDNSServer()
    
    http.HandleFunc("/resolve", server.HandleResolve)
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    
    fmt.Println("HTTP DNS Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}

// 客户端使用示例
func queryHTTPDNS(domain string) ([]string, error) {
    resp, err := http.Get(fmt.Sprintf("http://httpdns.example.com:8080/resolve?domain=%s", domain))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var record DNSRecord
    if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
        return nil, err
    }
    
    return record.IPs, nil
}
```

---

## 🔗 相关资源

- [CSDN原文：字节跳动蚂蚁金服百度SRE社招面经](https://blog.csdn.net/CCIEHL/article/details/104151990)
- [LeetCode算法练习](https://leetcode.cn/)
- [B+树原理详解](https://en.wikipedia.org/wiki/B%2B_tree)

## 📝 复习要点

1. **逻辑思维题重在思路**，不要急于求解
2. **算法题要掌握多种语言实现**
3. **装饰器等高级特性要理解原理**
4. **故障排查要有系统化的方法论**
5. **工具使用要了解原理和适用场景**
