package main

import (
	"testing"
)

func TestLc0890(t *testing.T) {
	// ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
	findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb")
}

func findAndReplacePattern(words []string, pattern string) []string {
	var check = func(s string) bool {
		// 双向映射
		fromTo, toFrom := make(map[byte]byte), make(map[byte]byte)
		for i := range s {
			from, to := s[i]-'a', pattern[i]-'a'
			if _, ok := fromTo[from]; ok {
				// 存在映射 反向也需要存在
				if _, ok = toFrom[to]; !ok || toFrom[to] != from {
					return false
				}
			} else {
				// 反向也不能存在
				if _, ok = toFrom[to]; ok {
					return false
				}
				fromTo[from] = to
				toFrom[to] = from
			}
		}
		return true
	}
	n := len(words)
	for i := 0; i < n; i++ {
		if check(words[i]) {
			continue
		}
		n--
		words[i], words[n] = words[n], words[i]
		i--
	}
	return words[:n]
}
