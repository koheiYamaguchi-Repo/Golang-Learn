package main

import "fmt"

type Data interface {
	Initial(name string, data []int)
	PrintData()
}
type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

func main() {
	var ob Mydata = Mydata{}
	ob.Initial("Sachiko", []int{56, 66, 77})
	ob.PrintData()
}
