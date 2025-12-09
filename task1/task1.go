package main

import (
	"fmt"
	"math"
	"sort"
)

// 只出现一次的数字
func findOnceNumber() {
	num := []int{1, 5, 5, 4, 4, 2, 1}
	numMap := make(map[int]int)
	for _, val := range num {
		numMap[val] += 1
	}
	for k, val := range numMap {
		if val == 1 {
			fmt.Println("Found number:", k)
		}
	}
}

// 是否回文数
func isPalindromeNumber() {
	num := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	numLen := int(math.Floor(float64(len(num)) / 2))
	count := 0

	for i := 0; i < numLen; i++ {
		if num[i] == num[len(num)-1-i] {
			count++
		}
	}
	if count == numLen {
		fmt.Println("是回文数:", num)
	} else {
		fmt.Println("is not PalindromeNumber:", num)
	}
}

// 有效括号
func validSymbol(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []rune
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == mp[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(sts []string) string {
	if len(sts) == 0 {
		return ""
	}
	if len(sts) == 1 {
		return sts[0]
	}
	prefix := sts[0]
	for i := 1; i < len(sts); i++ {
		prefix = commonPrefix(prefix, sts[i])
		if prefix == "" {
			break
		}
	}

	return prefix
}
func commonPrefix(str1 string, str2 string) string {
	minLen := len(str1)
	if len(str2) < minLen {
		minLen = len(str2)
	}

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:minLen]
}

func plusOne(nums []int) []int {

	nuLen := len(nums) - 1
	for i := nuLen; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}
	result := make([]int, len(nums)+1)
	result[0] = 1
	return result
}

func removeDuplicates1(nums []int) int {
	var result []int
	result = append(result, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != result[len(result)-1] {
			result = append(result, nums[i])
		}
	}
	return len(result)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[i-1] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 两数相加
func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}

func twoSumV2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, val := range nums {
		result := target - val
		if index, ok := numMap[result]; ok {
			return []int{index, i}
		}
		numMap[val] = i
	}
	return []int{}
}

func twoSumV3(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		if index, ok := m[val]; ok {
			return []int{index, i}
		}
		m[target-val] = i
	}
	return []int{}
}

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	current := intervals[0]
	var result [][]int
	for i := 1; i <= len(intervals); i++ {
		if i == len(intervals) {
			result = append(result, current)
			continue
		}
		next := intervals[i]
		if current[1] >= next[0] {
			current = []int{current[0], next[1]}
		} else {
			result = append(result, current)
			current = next
		}
	}
	return result
}

func mergeV2(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	current := intervals[0]
	var result [][]int
	for i := 1; i <= len(intervals); i++ {
		if i == len(intervals) {
			result = append(result, current)
			continue
		}
		next := intervals[i]
		if current[1] >= next[0] {
			current = []int{current[0], merMax(current[1], next[1])}
		} else {
			result = append(result, current)
			current = next
		}
	}
	return result
}

func mergeV3(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		next := result[len(result)-1]
		if interval[0] <= next[1] {
			next[1] = merMax(next[1], interval[0])
		} else {
			result = append(result, interval)
		}
	}
	return result
}

func merMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(validSymbol("()[]{}"))
	//fmt.Println(validSymbol("()]{}"))
	//fmt.Println(validSymbol("([])"))
	// 测试用例
	//testCases := [][]string{
	//	{"flower", "flow", "flight"},
	//	{"dog", "racecar", "car"},
	//}
	//
	//fmt.Println("=== 最长公共前缀 - 水平扫描法 ===")
	//for i, strs := range testCases {
	//	result := longestCommonPrefix(strs)
	//	fmt.Printf("测试 %d: %v -> %q\n", i+1, strs, result)
	//}
	//num := []int{2, 5, 4, 3, 5}
	//fmt.Println(twoSumV3(num, 10))
	intervals := [][]int{{1, 3}, {3, 6}, {7, 10}, {12, 18}, {15, 20}}
	fmt.Println(mergeV3(intervals))
}
