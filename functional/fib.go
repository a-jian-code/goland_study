package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 闭包实现斐波那契数列
// go语言闭包更为自然，不需要修饰如何访问自由变量，python中需要nolocal修饰
// 没有lambda表达式，但是有匿名函数
func fibonacci() intGen {
	last, now := 0, 1
	return func() int {
		last, now = now, last+now
		return last
	}
}

// 函数实习接口
type intGen func() int

func (ig intGen) Read(p []byte) (n int, err error) {
	next := ig()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	f := fibonacci()
	printFileContents(f)
}
