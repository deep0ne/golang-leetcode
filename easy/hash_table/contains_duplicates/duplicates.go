// https://leetcode.com/problems/contains-duplicate-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	hashTable := make(map[int]int)
	for idx, num := range nums {
		if hashIdx, ok := hashTable[num]; ok {
			if int(math.Abs(float64(idx)-float64(hashIdx))) <= k {
				return true
			}
		}
		hashTable[num] = idx
	}
	return false
}
