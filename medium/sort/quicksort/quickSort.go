// https://leetcode.com/problems/sort-an-array/description/

package main

import "fmt"

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	index := left
	for left <= right-1 {
		if nums[left] < pivot {
			nums[left], nums[index] = nums[index], nums[left]
			index++
		}
		left++
	}
	nums[right], nums[index] = nums[index], nums[right]
	return index
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	part := partition(nums, left, right)
	quickSort(nums, left, part-1)
	quickSort(nums, part+1, right)
}

func sortArray(nums []int) []int {
	l, r := 0, len(nums)-1
	quickSort(nums, l, r)
	return nums
}

func main() {
	s := []int{5, 2, 3, 1}
	fmt.Println(sortArray(s))
}
