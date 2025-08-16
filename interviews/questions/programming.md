# 🐹 编程技术面试题

> 适合：开发工程师、SRE工程师、技术面试
> 难度：⭐⭐⭐⭐ (中级-高级)

## 📋 Go语言进阶

### 1. 并发编程

#### Goroutine管理
```go
// Worker Pool模式
type Collector struct {
    workers    int
    jobQueue   chan Job
    resultChan chan Result
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

func (c *Collector) Start() {
    for i := 0; i < c.workers; i++ {
        c.wg.Add(1)
        go c.worker()
    }
}

func (c *Collector) worker() {
    defer c.wg.Done()
    for {
        select {
        case job := <-c.jobQueue:
            result := c.processJob(job)
            c.resultChan <- result
        case <-c.ctx.Done():
            return
        }
    }
}
```

#### 防止Goroutine泄漏
- 使用context控制生命周期
- 设置超时和取消机制
- 监控goroutine数量
- 优雅关闭和资源清理

### 2. 内存管理

#### 性能优化策略
```go
// 使用对象池减少GC压力
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

// 及时释放大对象
func processData(data []byte) {
    defer func() {
        data = nil // 帮助GC回收
    }()
    // 处理逻辑
}

// 使用context控制超时
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

#### 诊断工具
- **pprof**: CPU和内存分析
- **go tool trace**: 运行时追踪
- **runtime.ReadMemStats**: 内存统计

### 3. 网络编程

#### 高性能HTTP服务
```go
type Server struct {
    router *mux.Router
    server *http.Server
    config *Config
}

func (s *Server) Start() error {
    s.server = &http.Server{
        Addr:         s.config.Addr,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        IdleTimeout:  120 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    
    s.server.SetKeepAlivesEnabled(true)
    return s.server.ListenAndServe()
}
```

### 4. 错误处理

#### 结构化错误
```go
type AppError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Cause   error  `json:"cause,omitempty"`
    Stack   string `json:"stack,omitempty"`
}

func (e *AppError) Error() string {
    return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Cause)
}

// 错误包装
func wrapError(err error, message string) error {
    return &AppError{
        Code:    500,
        Message: message,
        Cause:   err,
        Stack:   string(debug.Stack()),
    }
}
```

## 📋 Python进阶

### 5. 迭代器和生成器

#### 迭代器实现
```python
class Counter:
    def __init__(self, max_count):
        self.max_count = max_count
        self.count = 0
    
    def __iter__(self):
        return self
    
    def __next__(self):
        if self.count < self.max_count:
            self.count += 1
            return self.count
        raise StopIteration
```

#### 生成器实现
```python
# 生成器函数
def counter_generator(max_count):
    count = 0
    while count < max_count:
        count += 1
        yield count

# 生成器表达式
squares = (x*x for x in range(5))
```

#### 主要区别
| 特性 | 迭代器 | 生成器 |
|------|--------|--------|
| 定义方式 | 实现__iter__和__next__ | 使用yield关键字 |
| 内存使用 | 需要维护状态 | 自动维护状态 |
| 代码复杂度 | 较复杂 | 简洁 |

### 6. 内存模型

#### is vs == 区别
```python
# == 比较值
list1 = [1, 2, 3]
list2 = [1, 2, 3]
print(list1 == list2)  # True，值相等
print(list1 is list2)  # False，不同对象

# is 比较身份（内存地址）
x = 256
y = 256
print(x is y)  # True，小整数缓存

# 特殊情况
- 小整数缓存：-5到256
- 小字符串缓存：简单字符串
- 单例对象：None、True、False
```

#### 深拷贝和浅拷贝
```python
import copy

# 浅拷贝
list1 = [[1, 2], [3, 4]]
list2 = copy.copy(list1)
list2[0][0] = 'X'  # 影响原列表

# 深拷贝
list3 = copy.deepcopy(list1)
list3[0][0] = 'Y'  # 不影响原列表
```

### 7. 垃圾回收机制

#### 引用计数
- 每个对象维护引用计数
- 引用计数为0时立即回收
- 无法处理循环引用

#### 标记清除
- 找到所有可达对象
- 清除不可达对象
- 处理循环引用

#### 分代回收
- 新对象放在第0代
- 存活对象提升到下一代
- 不同代使用不同回收策略

## 📋 算法和数据结构

### 8. 时间复杂度分析

```python
# O(1) - 常数时间
def get_first(arr):
    return arr[0] if arr else None

# O(n) - 线性时间
def linear_search(arr, target):
    for i, val in enumerate(arr):
        if val == target:
            return i
    return -1

# O(log n) - 对数时间
def binary_search(arr, target):
    left, right = 0, len(arr) - 1
    while left <= right:
        mid = (left + right) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1
```

### 9. 数据结构选择

#### 列表 vs 集合
```python
# 列表查找 - O(n)
if item in my_list:
    pass

# 集合查找 - O(1)
if item in my_set:
    pass

# 哈希表原理
- 通过哈希函数计算索引
- 直接定位到存储位置
- 平均时间复杂度O(1)
```

### 10. 系统调用

#### 文件操作
```python
# 上下文管理器
with open('file.txt', 'r') as f:
    content = f.read()

# 等价于
try:
    f = open('file.txt', 'r')
    content = f.read()
finally:
    f.close()
```

#### 进程和线程
```python
# 多进程
from multiprocessing import Process

def worker():
    print("Worker process")

p = Process(target=worker)
p.start()
p.join()

# 多线程
from threading import Thread

def worker():
    print("Worker thread")

t = Thread(target=worker)
t.start()
t.join()
```

## 🔗 相关资源

- [Go官方文档](https://golang.org/doc/)
- [Python官方文档](https://docs.python.org/3/)
- [算法可视化](https://visualgo.net/)

## 📝 复习要点

1. **掌握Go并发编程模型**
2. **理解Python内存管理机制**
3. **学会性能分析和优化**
4. **熟悉算法和数据结构**
5. **实践系统编程技巧**
