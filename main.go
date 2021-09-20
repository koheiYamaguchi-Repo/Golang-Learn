package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	var s1 = Mydata{"test1", []int{10, 30, 49}}
	fmt.Println(s1)
	rev(&s1)
	fmt.Println(s1)
}

func rev(md *Mydata) {
	od := md.Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
