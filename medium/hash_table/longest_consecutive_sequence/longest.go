// https://leetcode.com/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-interview-150
package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestConsecutive(nums []int) int {
	hashTable := make(map[int]struct{})
	for _, num := range nums {
		hashTable[num] = struct{}{}
	}
	count := 0
	for key := range hashTable {
		if _, ok := hashTable[key-1]; !ok {
			sequence := 1
			for {
				if _, ok := hashTable[key+sequence]; ok {
					sequence++
					continue
				}
				count = max(count, sequence)
				break
			}
		}
	}
	return count
}
