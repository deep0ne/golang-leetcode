package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	hashTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := hashTable[target-nums[i]]; ok {
			result = append(result, val, i)
		} else {
			hashTable[nums[i]] = i
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // must be [0, 1]
}
