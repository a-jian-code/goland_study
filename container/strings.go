package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网！"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //获取字符数量

	//将字符串传唤为aascii码，英文占一个字节，中文占三字节
	bytes := []byte(s) //[]byte获取字节
	fmt.Println(bytes, "-----")
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	//转换为rune对象遍历字符串可以拿到每一个下标的值，相当于char
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}

}
