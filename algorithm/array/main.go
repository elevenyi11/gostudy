package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

func main() {
	result := twoSum([]int{0, 1, 2, 5, 7, 11, 15}, 9)
	if result == nil {
		fmt.Println("not found")
	} else {
		for _, v := range result {
			fmt.Println(v)
		}
	}

	threeNums := []int{-1, 0, 1, 2, -1, -4}
	threeNumResult := threeSum(threeNums)
	if threeNumResult == nil {
		fmt.Println("not found threeNumResult")
	} else {
		for _, v := range threeNumResult {
			fmt.Println(v)
		}
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		sub := target - v
		if i, ok := m[sub]; ok {
			return []int{k, i}
		} else {
			m[v] = k
		}
	}
	return nil
}

// 用两层循环
func threeSum(nums []int) [][]int {
	numLen := len(nums)
	if numLen < 3 {
		return nil
	}
	result := make([][]int, 0, numLen)
	sort.Ints(nums)

	for k := 0; k < numLen-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		l, r := k+1, numLen-1
		if nums[k] > 0 || nums[k]+nums[l] > 0 {
			break
		}

		for l < r {
			if l > k+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-2 && nums[r] == nums[r+1] {
				r--
				continue
			}
			sum := nums[k] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[k], nums[l], nums[r]})
				l = l + 1
				r = r - 1
			} else if sum > 0 {
				r -= 1
			} else {
				l += 1
			}
		}
	}
	return result
}

func single() {
	one := sync.Once{}
	one.Do(func() {
		fmt.Println(runtime.NumGoroutine())
	})
}
