package main

import "fmt"

func sortArray(nums []int) []int {

	if len(nums) == 1 {
		return nums
	}

	left := sortArray(nums[:len(nums)/2])
	right := sortArray(nums[len(nums)/2:])

	numbers := make([]int, len(nums))

	l, r, k := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			numbers[k] = left[l]
			l++
		} else {
			numbers[k] = right[r]
			r++
		}
		k++
	}

	for l < len(left) {
		numbers[k] = left[l]
		l++
		k++
	}

	for r < len(right) {
		numbers[k] = right[r]
		r++
		k++
	}

	return numbers

}

func main() {
	s := []int{5, 2, 3, 1}
	fmt.Println(sortArray(s))
}
