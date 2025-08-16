# 💻 技术基础面试题

> 适合：应届生、初级工程师、技术基础复习
> 难度：⭐⭐⭐ (初级-中级)

## 📋 网络协议基础

### 1. OSI七层参考模型，以及对应有哪些协议

**标准答案**:
```
OSI七层模型及协议：

1. 物理层（Physical Layer）
   - 功能：传输原始比特流
   - 协议：以太网、光纤、无线电

2. 数据链路层（Data Link Layer）
   - 功能：帧传输、错误检测
   - 协议：Ethernet、PPP、Wi-Fi

3. 网络层（Network Layer）
   - 功能：路由选择、逻辑寻址
   - 协议：IP、ICMP、ARP、OSPF、BGP

4. 传输层（Transport Layer）
   - 功能：端到端通信、可靠传输
   - 协议：TCP、UDP

5. 会话层（Session Layer）
   - 功能：建立、管理、终止会话
   - 协议：NetBIOS、SQL、NFS

6. 表示层（Presentation Layer）
   - 功能：数据加密、压缩、格式转换
   - 协议：SSL/TLS、JPEG、MPEG

7. 应用层（Application Layer）
   - 功能：网络服务接口
   - 协议：HTTP、HTTPS、FTP、SMTP、DNS
```

### 2. TCP协议详解

#### TCP四次挥手过程
```
客户端                    服务器
FIN_WAIT_1  ----FIN---->  CLOSE_WAIT
FIN_WAIT_2  <---ACK----   CLOSE_WAIT
TIME_WAIT   <---FIN----   LAST_ACK
CLOSED      ----ACK---->  CLOSED
```

#### TCP可靠性机制
- 序列号和确认号
- 超时重传
- 滑动窗口
- 拥塞控制
- 校验和
- 连接管理

### 3. HTTP状态码详解

**常见状态码**:
- 200 OK：请求成功
- 404 Not Found：资源不存在
- 500 Internal Server Error：服务器内部错误
- 502 Bad Gateway：网关错误
- 503 Service Unavailable：服务不可用
- 504 Gateway Timeout：网关超时

## 📋 操作系统基础

### 4. Linux系统管理

#### 磁盘管理
```bash
# 查看磁盘使用情况
df -h

# 查看目录大小
du -sh /var/log

# df vs du区别
# df: 文件系统级别，显示可用空间
# du: 文件级别，显示已用空间
```

#### 内存管理
```bash
# 查看内存使用
free -h

# Buffer vs Cache
# Buffer: 写缓冲，内存 → 磁盘
# Cache: 读缓存，磁盘 → 内存
```

## 📋 编程基础

### 5. 算法题：数组相邻相同项相加

**问题**: [1,1,1,2,2,2,3,8,1,1] → [3,6,3,8,2]

**Go解法**:
```go
func mergeAdjacentSame(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    
    result := []int{}
    current := nums[0]
    count := 1
    
    for i := 1; i < len(nums); i++ {
        if nums[i] == current {
            count++
        } else {
            result = append(result, current*count)
            current = nums[i]
            count = 1
        }
    }
    
    result = append(result, current*count)
    return result
}
```

## 🔗 相关资源

- [计算机网络基础](https://www.runoob.com/tcpip/tcpip-tutorial.html)
- [Linux系统管理](https://www.runoob.com/linux/linux-tutorial.html)
- [算法题解](https://leetcode.cn/)

## 📝 复习要点

1. **掌握计算机网络基础概念**
2. **理解TCP/IP协议栈**
3. **熟悉Linux系统管理命令**
4. **学会基础算法和数据结构**
