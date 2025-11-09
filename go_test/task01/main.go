package main

import (
	"fmt"
	"sort"
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
// singleNumber 接收一个整型切片，返回只出现一次的元素；若不存在则返回 0
func singleNumber(nums []int) int {
	// 创建 map 保存每个数字的出现次数
	m := make(map[int]int)

	// 遍历输入数组，统计每个数字的出现频次
	for _, num := range nums {
		// 当前数字出现次数加一
		m[num]++
	}

	// 遍历统计结果，寻找只出现一次的数字
	for key, freq := range m {
		// 如果当前数字只出现了一次，则输出并返回该数字
		if freq == 1 {
			fmt.Println("只出现了一次的元素: ", key)
			return key
		}
	}
	// 遍历结束仍未找到，说明不存在只出现一次的数字
	fmt.Println("只出现一次的元素: 没有找到")
	// 返回默认值 0
	return 0
}

// 任务2
// 题目：判断一个整数是否是回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

// isPalindrome 判断整数 x 是否是回文数，是则返回 true，否则 false
func isPalindrome(x int) bool {
	// 将整数转换为字符串形式，便于比较
	str := strconv.Itoa(x)
	// 只需要比较字符串前半部分与后半部分是否对称相等
	for i := 0; i < len(str)/2; i++ {
		// 如果首尾对应字符不相等，则不是回文数
		if str[i] != str[len(str)-i-1] {
			fmt.Println("不是回文数")
			return false
		}
	}
	// 若循环结束没有发现不相等的情况，则是回文数
	fmt.Println("是回文数")
	return true
}

// 任务3
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
// isValid 判断由括号组成的字符串是否有效（成对且顺序正确）
func isValid(s string) bool {
	// 使用切片模拟栈结构保存未匹配的左括号
	var stack []rune

	// 定义括号的映射关系：右括号 -> 左括号
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// 根据当前字符类型分别处理
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
// longestCommonPrefixV2 返回字符串数组中的最长公共前缀，如果不存在则返回空串
func longestCommonPrefixV2(strs []string) string {
	// 如果输入为空直接返回空字符串
	if len(strs) == 0 {
		return ""
	}

	// 遍历第一个字符串的每个字符位置
	for i := 0; i < len(strs[0]); i++ {
		// 取出第一个字符串当前位置的字符作为比较基准
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

// plusOne 将一个代表整数的数组加一，并返回进位处理后的新数组
func plusOne(digits []int) []int {
	// 记录数组长度，便于从后往前遍历
	n := len(digits)
	// 从最低位开始处理进位
	for i := n - 1; i >= 0; i-- {
		// 如果当前位小于 9，加一后不会产生进位，可直接返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 当前位是 9，加一后变成 0，并继续向前进位
		digits[i] = 0
	}
	// 所有位都是 9 的情况，需要在最前面补 1
	return append([]int{1}, digits...)
}

// 任务6
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。

// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。

// removeDuplicates 原地删除升序数组中的重复元素，返回去重后的长度
func removeDuplicates(nums []int) int {
	// 特殊情况：空数组直接返回 0
	if len(nums) == 0 {
		return 0
	}

	// 慢指针指向已确认的唯一元素末尾
	slow := 0
	// 快指针遍历数组寻找新元素
	for fast := 1; fast < len(nums); fast++ {
		// 遇到与慢指针不同的元素，说明找到了新的唯一值
		if nums[fast] != nums[slow] {
			// 先移动慢指针位置
			slow++
			// 将新的唯一值放到慢指针位置，实现原地覆盖
			nums[slow] = nums[fast]
		}
	}
	// 打印去重后的数组内容
	fmt.Println(nums[:slow+1])
	// 打印去重后数组的长度
	fmt.Println(slow + 1)
	// 返回去重后长度
	return slow + 1
}

// 任务7
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
// merge 合并所有重叠区间，并返回一个覆盖所有区间的结果集合
func merge(intervals [][]int) [][]int {
	// 如果区间列表为空或只有一个区间，直接返回即可
	if len(intervals) < 2 {
		return intervals
	}

	// 先对所有区间按起点从小到大排序，保证后续合并时可以线性扫描
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用于保存合并后的结果
	var merged [][]int
	// 初始化当前正在合并的区间为排序后的第一个区间
	current := intervals[0]

	// 遍历剩余区间并尝试与当前区间合并
	for i := 1; i < len(intervals); i++ {
		// 取出下一个区间与当前区间进行比较
		next := intervals[i]

		// 如果下一个区间与当前区间有重叠（起点 <= 当前区间的终点）
		if next[0] <= current[1] {
			// 更新当前区间的终点为两者中的较大值，实现合并
			if next[1] > current[1] {
				current[1] = next[1]
			}
		} else {
			// 没有重叠，则将当前区间保存到结果中，并开始新的区间
			merged = append(merged, current)
			current = next
		}
	}

	// 循环结束后别忘了把最后一个区间加入结果
	merged = append(merged, current)

	// 返回最终的合并结果
	return merged
}

// 任务8
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

// 你可以按任意顺序返回答案。

// twoSum2 在数组中寻找任意两个数，使其和为 target，返回这两个数的下标
func twoSum2(nums []int, target int) []int {
	// 使用哈希表记录「值 -> 下标」，单次遍历即可找到答案
	seen := make(map[int]int)

	// 遍历每个元素并查找与之匹配的目标值
	for i, num := range nums {
		// 计算当前元素需要配对的目标值
		need := target - num

		// 如果之前已经遇到过需要的数，直接返回两个下标
		if idx, ok := seen[need]; ok {
			return []int{idx, i}
		}

		// 否则把当前数记录下来，供后面的元素匹配
		seen[num] = i
	}

	// 题目保证一定有答案，此处返回空切片作为兜底
	return []int{}
}

// twoSum 使用哈希表在 O(n) 时间内返回两数之和等于 target 的下标
func twoSum(nums []int, target int) []int {
	// 使用哈希表存储已经遍历过的数字及其下标
	prevNums := map[int]int{}
	// 遍历数组中的每个数字
	for i, num := range nums {
		// 计算与当前数字配对所需的目标值
		targetNum := target - num
		// 查看目标值是否已经出现过
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			// 若存在，直接返回两个数字的下标
			return []int{targetNumIndex, i}
		} else {
			// 若不存在，将当前数字及下标记录下来
			prevNums[num] = i
		}
	}
	// 题目保证一定有答案，兜底返回空切片
	return []int{}
}

func main() {
	// 测试只出现一次的数字
	singleNumber([]int{2, 2, 1, 1, 3})

	// 测试回文数判断
	isPalindrome(1234554321)

	// 测试括号匹配有效性
	if isValid("({[]})") {
		// 当括号有效时输出提示
		fmt.Println("是有效括号")
	} else {
		// 当括号无效时输出提示
		fmt.Println("不是有效括号")
	}

	// 打印最长公共前缀结果
	fmt.Println(longestCommonPrefixV2([]string{"flower", "flow", "flight"}))

	// 打印加一后的数组
	fmt.Println(plusOne([]int{9, 9, 9, 9}))

	// 执行数组去重示例
	removeDuplicates([]int{1, 2, 3, 3, 4, 5, 5, 5, 6, 6, 6, 7, 7, 8})

	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	// 打印两数之和示例的下标
	fmt.Println(twoSum([]int{2, 8, 7, 11, 15}, 9))
}
