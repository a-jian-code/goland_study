package main

import (
	"fmt"
)

//if 判断语句
//func main() {
//	const filename = "a.txt"
//	if contents, err := ioutil.ReadFile(filename); err == nil {
//		fmt.Println(string(contents))
//	} else {
//		fmt.Println(err)
//	}
//}

// switch 语句
func grade(score int) string {
	g := ""
	switch {
	//抛出异常
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	println(grade(40))
}
