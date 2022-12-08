package main

import "fmt"

func main() {
	//定义方式创建map
	m := map[string]string{
		"name": "goland",
		"age":  "18",
	}

	//make方法创建
	m1 := make(map[string]int)
	fmt.Println(m, m1, m["name"])
	//遍历map
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	//获取value
	fmt.Println("Getting values")
	//键不存在时获取到空值，不会报错
	value, exist := m["name"]
	fmt.Println(value, exist)
	if value1, exist1 := m["naem"]; exist1 {
		fmt.Println(value1)
	} else {
		fmt.Println("key does not exist")
	}
	//删除value
	fmt.Println("Deleting values")
	delete(m, "name")
	name, exi := m["name"]
	fmt.Println(name, exi)

}
