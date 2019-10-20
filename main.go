package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	table := map[byte][]byte{
		'2' : []byte{'a', 'b', 'c'},
		'3' : []byte{'d', 'e', 'f'},
		'4' : []byte{'g', 'h', 'i'},
		'5' : []byte{'j', 'k', 'l'},
		'6' : []byte{'m', 'n', 'o'},
		'7' : []byte{'p', 'q', 'r', 's'},
		'8' : []byte{'t', 'u', 'v'},
		'9' : []byte{'w', 'x', 'y', 'z'},
	}
	dig := [][]byte{}
	ans := []string{}
	for i := range digits {
		dig = append(dig, table[digits[i]])
	}
	for i := 0; i < len(dig[0]); i++ {
		ans = getAns(dig, 1, []byte{dig[0][i]}, ans)
	}
	return ans[:len(ans)-1]
}

func getAns(digits [][]byte, depth int, tmp []byte, ans []string) []string {
	if depth >= len(digits) {
		return append(ans, string(tmp))
	}
	for i := 0; i < len(digits[depth]); i++ {
		ans = getAns(digits, depth+1, append(tmp, digits[depth][i]), ans)
	}
	return ans
}