// https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := make([]string, 0)
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			sb.WriteByte(s[i])
		} else {
			if sb.Len() == 0 {
				continue
			}
			word := sb.String()
			if len(word) > 0 && word[len(word)-1] != ' ' {
				words = append(words, word)
				sb.Reset()
			}
		}
	}
	if sb.Len() != 0 {
		words = append(words, sb.String())
	}
	sb.Reset()
	for i := len(words) - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		if i != 0 {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}

func main() {
	s := "  hello world  "
	fmt.Println(reverseWords(s)) // must be "world hello"
}
