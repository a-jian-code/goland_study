package main

import (
	"fmt"
	"goland/retriever/mock"
	"goland/retriever/real"
	"time"
)

// interface{}代表任何类型
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}
func post(poster Poster) {
	poster.Post(url, map[string]string{"name": "lj", "course": "goland"})
}

// 继承接口，并且可以增加自己的接口，但是方法名不能与被继承接口重名
type RetrieverPost interface {
	Retriever
	Poster
}

func session(s RetrieverPost) string {
	s.Post(url, map[string]string{
		"contents": "another facked imooc.com",
	})
	return s.Get(url)
}
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(">%T  %v\n", r, r)
	fmt.Printf(">Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
func main() {
	var r Retriever
	retriever := mock.Retriever{"This is a fake imooc.com"}
	r = &mock.Retriever{"This is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)

	//type assertion
	//realRetriever := r.(*real.Retriever)
	//fmt.Println(realRetriever.TimeOut)

	//if mockRetriever, ok := r.(mock.Retriever); ok {
	//	fmt.Println(mockRetriever.Contents)
	//} else {
	//	panic("This is not a mock retriever")
	//}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
