package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ============ LeetCode前30题精选Go实现 ============

// 1. 两数之和 (Two Sum)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

// 2. 两数相加 (Add Two Numbers)
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0
	
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	
	return dummy.Next
}

// 3. 无重复字符的最长子串 (已在主文档中实现)

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	
	for left <= right {
		partitionX := (left + right) / 2
		partitionY := (m+n+1)/2 - partitionX
		
		maxLeftX := getMax(nums1, partitionX-1)
		minRightX := getMin(nums1, partitionX)
		
		maxLeftY := getMax(nums2, partitionY-1)
		minRightY := getMin(nums2, partitionY)
		
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			right = partitionX - 1
		} else {
			left = partitionX + 1
		}
	}
	return 0.0
}

func getMax(nums []int, index int) int {
	if index < 0 {
		return -1000000
	}
	return nums[index]
}

func getMin(nums []int, index int) int {
	if index >= len(nums) {
		return 1000000
	}
	return nums[index]
}

// 5. 最长回文子串
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	
	start, maxLen := 0, 1
	
	for i := 0; i < len(s); i++ {
		// 奇数长度回文
		len1 := expandAroundCenter(s, i, i)
		// 偶数长度回文
		len2 := expandAroundCenter(s, i, i+1)
		
		currentMax := max(len1, len2)
		if currentMax > maxLen {
			maxLen = currentMax
			start = i - (currentMax-1)/2
		}
	}
	
	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 7. 整数反转
func reverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		
		// 检查溢出
		if result > (1<<31-1)/10 || (result == (1<<31-1)/10 && digit > 7) {
			return 0
		}
		if result < (-1<<31)/10 || (result == (-1<<31)/10 && digit < -8) {
			return 0
		}
		
		result = result*10 + digit
	}
	return result
}

// 8. 字符串转换整数 (atoi)
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	
	sign := 1
	index := 0
	
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		index++
	}
	
	result := 0
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		digit := int(s[index] - '0')
		
		// 检查溢出
		if result > (1<<31-1)/10 || (result == (1<<31-1)/10 && digit > 7) {
			if sign == 1 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}
		
		result = result*10 + digit
		index++
	}
	
	return sign * result
}

// 9. 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	
	return x == reversed || x == reversed/10
}

// 11. 盛最多水的容器
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	
	for left < right {
		water := min(height[left], height[right]) * (right - left)
		maxWater = max(maxWater, water)
		
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	
	return maxWater
}

// 13. 罗马数字转整数
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}
	
	return result
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	
	return prefix
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	
	// 排序
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复元素
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

// 20. 有效的括号
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	
	return len(stack) == 0
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	
	return dummy.Next
}

// 26. 删除有序数组中的重复项
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

// 工具函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ============ 测试函数 ============
func testAlgorithms() {
	// 测试两数之和
	fmt.Println("1. 两数之和:")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("输入: %v, 目标: %d, 输出: %v\n\n", nums, target, result)
	
	// 测试无重复字符最长子串
	fmt.Println("3. 无重复字符最长子串:")
	s := "abcabcbb"
	length := lengthOfLongestSubstring(s)
	fmt.Printf("输入: %s, 输出: %d\n\n", s, length)
	
	// 测试回文数
	fmt.Println("9. 回文数:")
	x := 121
	isPalin := isPalindrome(x)
	fmt.Printf("输入: %d, 输出: %t\n\n", x, isPalin)
	
	// 测试有效括号
	fmt.Println("20. 有效括号:")
	brackets := "()[]{}"
	isValidBrackets := isValid(brackets)
	fmt.Printf("输入: %s, 输出: %t\n\n", brackets, isValidBrackets)
}

func main() {
	fmt.Println("============ 算法题练习 ============")
	testAlgorithms()
} 