package main

import (
	"strings"
)

// Function finds the first unique character from given string.
// Returns rune and ok boolean
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

// Function actually takes slice of strings and returns
// slice of the first unique characters in each sting of given slice.
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

// Exported fuction. Returns the first unique character and ok boolean from the slice of
// unique characters from each substing separated by whitespace of given string.
func UniqOfUniqsChar(s string) (rune, bool) {
	arrStr := strings.Split(s, " ")
	arrUniqChars := uniqCharsFromArr(arrStr)
	return uniqChar(string(arrUniqChars))
}
