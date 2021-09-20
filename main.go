package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func (md Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

func main() {
	taro := new(Mydata)
	taro.Name = "hanako"
	taro.Data = []int{30, 49, 29}
	taro.PrintData()
}
