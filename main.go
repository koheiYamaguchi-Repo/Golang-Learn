package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	var s1 = Mydata{"test1", []int{10, 30, 49}}
	var s2 = Mydata{"test2", []int{39, 94, 22}}
	fmt.Println(s1, s2)
}
