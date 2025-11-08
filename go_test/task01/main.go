package main

import (
	"fmt"
	"strconv"
)

/*
*
// 任务1
控制流程
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

回文数

考察：数字操作、条件判断
*/
func singleNumber(nums []int) int {

	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	for key, freq := range m {
		if freq == 1 {
			fmt.Println("只出现了一次的元素: ", key)
			return key
		}
	}
	fmt.Println("只出现一次的元素: 没有找到")
	return 0
}

// 任务2
// 题目：判断一个整数是否是回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			fmt.Println("不是回文数")
			return false
		}
	}
	fmt.Println("是回文数")
	return true
}

// 任务3
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	// 使用切片模拟栈
	var stack []rune

	// 定义括号的映射关系：右括号 -> 左括号
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 如果是左括号，压入栈
			stack = append(stack, char)
		case ')', '}', ']':
			// 如果是右括号
			if len(stack) == 0 {
				// 栈为空，说明没有匹配的左括号
				return false
			}
			// 弹出栈顶元素
			top := stack[len(stack)-1]
			// 移除最后一个元素
			stack = stack[:len(stack)-1]

			// 检查是否匹配
			if top != pairs[char] {
				return false
			}
		}
	}

	// 最后栈必须为空才是有效括号
	return len(stack) == 0
}

// 任务4
// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 直扫描法（更高效）
func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 遍历第一个字符串的每个字符位置
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		// 检查其他所有字符串在位置i的字符是否相同
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够，或者字符不匹配
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	// 如果第一个字符串的所有字符都匹配，则返回整个第一个字符串
	return strs[0]
}

// 任务5
// 基本值类型
// 加一

// 难度：简单

// 考察：数组操作、进位处理

// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

// 任务6
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。

// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	fmt.Println(nums[:slow+1])
	fmt.Println(slow + 1)
	return slow + 1
}

// 【任务7】
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {

	return intervals
}

// 任务8
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	prevNums := map[int]int{}
	for i, num := range nums {
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			return []int{targetNumIndex, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}

func main() {
	singleNumber([]int{2, 2, 1, 1, 3})

	isPalindrome(1234554321)

	if isValid("({[]})") {
		fmt.Println("是有效括号")
	} else {
		fmt.Println("不是有效括号")
	}

	fmt.Println(longestCommonPrefixV2([]string{"flower", "flow", "flight"}))

	fmt.Println(plusOne([]int{9, 9, 9, 9}))

	removeDuplicates([]int{1, 2, 3, 3, 4, 5, 5, 5, 6, 6, 6, 7, 7, 8})

	fmt.Println(twoSum([]int{2, 8, 7, 11, 15}, 9))
}
