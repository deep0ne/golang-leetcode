// https://leetcode.com/problems/majority-element/description/

package main

import "fmt"

func majorityElement(nums []int) int {
	hashTable := make(map[int]int)
	var res int
	for _, num := range nums {
		hashTable[num]++
		if count, _ := hashTable[num]; count > len(nums)/2 {
			res = num
		}
	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3})) // must be 3
}
