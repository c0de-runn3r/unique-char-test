package main

import (
	"strings"
)

func uniqChar(s string) (rune, bool) {
	m := make(map[rune]int)
	for _, v := range s {
		_, ok := m[v]
		if !ok {
			m[v] = 0
		}
		m[v]++
	}
	for _, v := range s {
		x := m[v]
		if x == 1 {
			return v, true
		}
	}
	return -1, false
}

func uniqCharsFromArr(arr []string) []rune {
	charArr := make([]rune, 0)
	for i := 0; i < len(arr); i++ {
		char, ok := uniqChar(arr[i])
		if ok {
			charArr = append(charArr, char)
		}
	}
	return charArr
}

func UniqOfUniqsChar(s string) (rune, bool) {
	arrStr := strings.Split(s, " ")
	arrUniqChars := uniqCharsFromArr(arrStr)
	return uniqChar(string(arrUniqChars))
}
