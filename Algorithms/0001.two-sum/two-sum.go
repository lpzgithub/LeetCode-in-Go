package problem0001

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(nums))

	// 通过 for 循环，获取b的序列号
	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			// ok 为 true
			// 说明在i之前，存在 nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i
		}

		// 把b和i的值，存入map
		index[b] = i
	}

	return nil
}

func twoSumSort(nums []int, target int) []int {
	sort.Slice(nums, func(i int, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})

	lenght := len(nums)
	for i, j := 0, lenght-1; i < j; {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return nil
}
