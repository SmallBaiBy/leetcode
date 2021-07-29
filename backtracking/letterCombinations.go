package backtracking

import (
	"fmt"
	"strings"
)

var tel = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

//Solution02 17. 电话号码的字母组合
func Solution02(digits string) []string {
	var result02 []string
	var path02 string
	split := strings.Split(digits, "")
	if len(split) == 0 {
		return nil
	}
	n := len(split)
	backtracking02(n, split, 0, path02, &result02)
	fmt.Println(result02)
	return result02
}

func backtracking02(n int, split []string, startIndex int, path02 string, result02 *[]string) {
	if len(path02) == n {
		*result02 = append(*result02, path02)
		return
	}
	layer := tel[split[startIndex]]
	layers := strings.Split(layer, "")
	for i := 0; i < len(layers); i++ {
		path02 += layers[i]
		backtracking02(n, split, startIndex+1, path02, result02)
		path02 = string(path02[:len(path02)-1])
	}
}
