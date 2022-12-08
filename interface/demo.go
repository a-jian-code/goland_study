package main

import "fmt"

// 定义开关接口
type Kaiguan interface {
	start(name string)
	stop(name string)
}

// 定义开关实现类
type cow_kaiguan_impl struct {
}
type xm_kaiguan_impl struct {
}

// 公牛开关实现类结构体实现interface里start方法
func (cow cow_kaiguan_impl) start(name string) {
	fmt.Println(name, "打开公牛开关")
	fmt.Println("公牛开关实现类，实现start方法")
}

// 公牛开关实现类结构体实现interface里stop方法
func (cow cow_kaiguan_impl) stop(name string) {
	fmt.Println(name, "关闭公牛开关")
	fmt.Println("公牛开关实现类，实现stop方法")
}

func (xm xm_kaiguan_impl) start(name string) {
	fmt.Println("打开小米开关")
	fmt.Println("小米开关实现类，实现start方法")
}

func (xm xm_kaiguan_impl) stop(name string) {
	fmt.Println("关闭小米开关")
	fmt.Println("小米开关实现类，实现stop方法")
}

type People struct {
	Name string
}

func (p *People) use_kaiguan(kaiguan Kaiguan) {
	kaiguan.start(p.Name)
	kaiguan.stop(p.Name)
}

func main() {
	people := People{"小明"}
	gn := cow_kaiguan_impl{}
	xm := xm_kaiguan_impl{}
	people.use_kaiguan(gn)
	fmt.Println("---------------------------")
	people.use_kaiguan(xm)
}
