package main

import "fmt"

// 寻找最长不含有重复字符的字串
// lastOccured[x]不存在，或者<start无需操作
// lastOccured[x]>=start  更新start
// 更新lastOccured[x],更新maxLength

// 英文版
func lengthOfNonRepeatingSubStr(s string) (int, string) {
	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0
	sub := ""
	for i, ch := range []byte(s) {

		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			sub = s[start : i+1]
		}
		lastOccured[ch] = i
	}
	return maxLength, sub
}

// 国际版
func lengthOfNonRepeatingSubStr1(s string) (int, string) {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	sub := ""
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			sub = s[start : i+1]
		}
		lastOccured[ch] = i
	}
	return maxLength, sub
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("aaaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcde"))
	fmt.Println(lengthOfNonRepeatingSubStr1("你好你abcde"))
}
