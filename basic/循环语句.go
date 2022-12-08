package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//1-100求和
//func main() {
//	sum := 0
//	for i := 1; i < 100; i++ {
//		sum += i
//	}
//	println(sum)
//}

// 十进制转换为二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// go里面没有while  用for循环可以实现
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环 for什么都不加
func forever() {
	for {
		fmt.Println("hhh")
	}
}

func main() {
	println(convertToBin(0))
	printFile("a.txt")
	//forever()
}
