# 🧮 高频算法题集合

> 30分钟刷完认证考试必考算法题，代码可直接运行

## 📋 题目清单

### 🎯 数组类（必考）
1. [两数之和](#1-两数之和) - HashMap O(n)
2. [三数之和](#2-三数之和) - 双指针 O(n²)
3. [删除重复元素](#3-删除重复元素) - 双指针 O(n)
4. [多数元素](#4-多数元素) - Boyer-Moore O(n)

### 🔗 链表类（高频）
5. [反转链表](#5-反转链表) - 迭代/递归 O(n)
6. [合并有序链表](#6-合并有序链表) - 双指针 O(n)
7. [两两交换节点](#7-两两交换节点) - 指针操作 O(n)
8. [环形链表检测](#8-环形链表检测) - 快慢指针 O(n)

### 📝 字符串类（常考）
9. [最长无重复子串](#9-最长无重复子串) - 滑动窗口 O(n)
10. [有效括号](#10-有效括号) - 栈 O(n)

---

## 🎯 数组类

### 1. 两数之和
**题目**: 给定数组和目标值，返回两数之和等于目标值的索引

```go
func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        if index, exists := hashMap[complement]; exists {
            return []int{index, i}
        }
        hashMap[num] = i
    }
    
    return nil
}

// 测试用例
func TestTwoSum() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Printf("输入: %v, 目标: %d\n", nums, target)
    fmt.Printf("输出: %v\n", result) // [0, 1]
}
```

### 2. 三数之和
**题目**: 找出数组中所有和为0的三元组

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    
    for i := 0; i < len(nums)-2; i++ {
        // 跳过重复元素
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        left, right := i+1, len(nums)-1
        
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            
            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                
                // 跳过重复元素
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                
                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }
    
    return result
}
```

### 3. 删除重复元素
**题目**: 原地删除排序数组中的重复项

```go
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
func TestRemoveDuplicates() {
    nums := []int{1, 1, 2, 2, 3, 4, 4}
    newLen := removeDuplicates(nums)
    fmt.Printf("新长度: %d, 数组: %v\n", newLen, nums[:newLen])
    // 输出: 新长度: 4, 数组: [1 2 3 4]
}
```

### 4. 多数元素
**题目**: 找出数组中出现次数超过n/2的元素

```go
func majorityElement(nums []int) int {
    candidate := nums[0]
    count := 1
    
    // Boyer-Moore投票算法
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
```

---

## 🔗 链表类

### 5. 反转链表
**题目**: 反转单链表

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

// 迭代方法
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    
    return prev
}

// 递归方法
func reverseListRecursive(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    newHead := reverseListRecursive(head.Next)
    head.Next.Next = head
    head.Next = nil
    
    return newHead
}
```

### 6. 合并有序链表
**题目**: 合并两个有序链表

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }
    
    // 连接剩余节点
    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }
    
    return dummy.Next
}
```

### 7. 两两交换节点
**题目**: 两两交换链表中的节点

```go
func swapPairs(head *ListNode) *ListNode {
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
```

### 8. 环形链表检测
**题目**: 检测链表是否有环

```go
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    
    slow := head
    fast := head.Next
    
    for slow != fast {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return true
}

// 找环的起始位置
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    slow, fast := head, head
    
    // 第一次相遇
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    
    if fast == nil || fast.Next == nil {
        return nil
    }
    
    // 找环的起始位置
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    
    return slow
}
```

---

## 📝 字符串类

### 9. 最长无重复子串
**题目**: 找出字符串中最长无重复字符的子串长度

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

// 测试
func TestLongestSubstring() {
    testCases := []string{"abcabcbb", "bbbbb", "pwwkew", ""}
    for _, s := range testCases {
        result := lengthOfLongestSubstring(s)
        fmt.Printf("输入: %s, 输出: %d\n", s, result)
    }
}
```

### 10. 有效括号
**题目**: 判断括号字符串是否有效

```go
func isValid(s string) bool {
    stack := []rune{}
    mapping := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    
    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            // 左括号入栈
            stack = append(stack, char)
        } else {
            // 右括号检查匹配
            if len(stack) == 0 {
                return false
            }
            
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            if top != mapping[char] {
                return false
            }
        }
    }
    
    return len(stack) == 0
}

// 测试
func TestIsValid() {
    testCases := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
    for _, s := range testCases {
        result := isValid(s)
        fmt.Printf("输入: %s, 输出: %t\n", s, result)
    }
}
```

---

## 🚀 完整测试代码

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println("=== 算法题测试 ===")
    
    // 测试两数之和
    fmt.Println("\n1. 两数之和:")
    TestTwoSum()
    
    // 测试删除重复元素
    fmt.Println("\n2. 删除重复元素:")
    TestRemoveDuplicates()
    
    // 测试最长无重复子串
    fmt.Println("\n3. 最长无重复子串:")
    TestLongestSubstring()
    
    // 测试有效括号
    fmt.Println("\n4. 有效括号:")
    TestIsValid()
    
    fmt.Println("\n=== 测试完成 ===")
}
```

---

## 💡 解题技巧总结

### 🎯 常用算法模式
1. **双指针**: 数组去重、三数之和、链表操作
2. **滑动窗口**: 最长子串、子数组问题
3. **哈希表**: 两数之和、字符统计
4. **栈**: 括号匹配、单调栈
5. **快慢指针**: 链表环检测、中点查找

### ⚡ 时间复杂度优化
- **O(n²) → O(n)**: 使用哈希表
- **O(n²) → O(n log n)**: 先排序再双指针
- **递归 → 迭代**: 避免栈溢出

### 🔍 边界情况检查
- 空数组/空链表
- 单元素情况
- 重复元素处理
- 整数溢出

---

**🎪 考试建议**: 
1. 先说思路再写代码
2. 考虑边界情况
3. 分析时空复杂度
4. 提供多种解法 