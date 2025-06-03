package main

import (
	"fmt"
	"sort"
)

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 两数之和
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

// 2. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

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

// 3. 删除重复元素
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

// 4. 多数元素
func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

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

// 5. 反转链表
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

// 6. 合并有序链表
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

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

// 7. 两两交换节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
	}

	return dummy.Next
}

// 8. 环形链表检测
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

// 9. 最长无重复子串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	charMap := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		if pos, exists := charMap[s[right]]; exists && pos >= left {
			left = pos + 1
		}

		charMap[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

// 10. 有效括号
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
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

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 创建链表的辅助函数
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}

	return head
}

// 打印链表的辅助函数
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// 测试函数
func testTwoSum() {
	fmt.Println("=== 1. 两数之和 ===")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("输入: %v, 目标: %d\n", nums, target)
	fmt.Printf("输出: %v\n\n", result)
}

func testThreeSum() {
	fmt.Println("=== 2. 三数之和 ===")
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Printf("输入: %v\n", nums)
	fmt.Printf("输出: %v\n\n", result)
}

func testRemoveDuplicates() {
	fmt.Println("=== 3. 删除重复元素 ===")
	nums := []int{1, 1, 2, 2, 3, 4, 4}
	fmt.Printf("原数组: %v\n", nums)
	newLen := removeDuplicates(nums)
	fmt.Printf("新长度: %d, 数组: %v\n\n", newLen, nums[:newLen])
}

func testMajorityElement() {
	fmt.Println("=== 4. 多数元素 ===")
	nums := []int{3, 2, 3}
	result := majorityElement(nums)
	fmt.Printf("输入: %v\n", nums)
	fmt.Printf("多数元素: %d\n\n", result)
}

func testReverseList() {
	fmt.Println("=== 5. 反转链表 ===")
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("原链表: ")
	printList(head)

	reversed := reverseList(head)
	fmt.Print("反转后: ")
	printList(reversed)
	fmt.Println()
}

func testMergeTwoLists() {
	fmt.Println("=== 6. 合并有序链表 ===")
	l1 := createList([]int{1, 2, 4})
	l2 := createList([]int{1, 3, 4})

	fmt.Print("链表1: ")
	printList(l1)
	fmt.Print("链表2: ")
	printList(l2)

	merged := mergeTwoLists(l1, l2)
	fmt.Print("合并后: ")
	printList(merged)
	fmt.Println()
}

func testSwapPairs() {
	fmt.Println("=== 7. 两两交换节点 ===")
	head := createList([]int{1, 2, 3, 4})
	fmt.Print("原链表: ")
	printList(head)

	swapped := swapPairs(head)
	fmt.Print("交换后: ")
	printList(swapped)
	fmt.Println()
}

func testHasCycle() {
	fmt.Println("=== 8. 环形链表检测 ===")
	// 创建无环链表
	head1 := createList([]int{3, 2, 0, -4})
	result1 := hasCycle(head1)
	fmt.Printf("无环链表检测结果: %t\n", result1)

	// 创建有环链表
	head2 := createList([]int{3, 2, 0, -4})
	// 手动创建环：让最后一个节点指向第二个节点
	current := head2
	var second *ListNode
	count := 0
	for current.Next != nil {
		if count == 1 {
			second = current
		}
		current = current.Next
		count++
	}
	current.Next = second // 创建环

	result2 := hasCycle(head2)
	fmt.Printf("有环链表检测结果: %t\n\n", result2)
}

func testLongestSubstring() {
	fmt.Println("=== 9. 最长无重复子串 ===")
	testCases := []string{"abcabcbb", "bbbbb", "pwwkew", ""}
	for _, s := range testCases {
		result := lengthOfLongestSubstring(s)
		fmt.Printf("输入: \"%s\", 输出: %d\n", s, result)
	}
	fmt.Println()
}

func testIsValid() {
	fmt.Println("=== 10. 有效括号 ===")
	testCases := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, s := range testCases {
		result := isValid(s)
		fmt.Printf("输入: \"%s\", 输出: %t\n", s, result)
	}
	fmt.Println()
}

func main() {
	fmt.Println("🧮 高频算法题测试")
	fmt.Println("==================")

	testTwoSum()
	testThreeSum()
	testRemoveDuplicates()
	testMajorityElement()
	testReverseList()
	testMergeTwoLists()
	testSwapPairs()
	testHasCycle()
	testLongestSubstring()
	testIsValid()

	fmt.Println("✅ 所有测试完成！")
	fmt.Println("💡 这些是面试中最常考的算法题，建议反复练习直到能够快速实现。")
}
