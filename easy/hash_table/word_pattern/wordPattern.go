// https://leetcode.com/problems/word-pattern/description/?envType=study-plan-v2&envId=top-interview-150
package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	hashTable1 := make(map[byte]string)
	hashTable2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if word, ok := hashTable1[pattern[i]]; ok {
			if words[i] != word {
				return false
			}
		}
		if letter, ok := hashTable2[words[i]]; ok {
			if pattern[i] != letter {
				return false
			}
		}
		hashTable1[pattern[i]] = words[i]
		hashTable2[words[i]] = pattern[i]
	}
	return true
}
