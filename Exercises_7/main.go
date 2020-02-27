package main

import (
	"fmt"
)

func main()  {
	Test1()
	Test2()
}

func Test1(){
	x := 42
	fmt.Println("Address: ", &x)
}

type person struct {
	firstName string
	lastName string
	age int
}

func (p *person) changeMe()  {
	p.age = 15
	(*p).firstName = "Nezuko"
}

func Test2(){
	p1 := person{
		firstName: "Jayson",
		lastName: "Tanjiro",
		age: 12,
	}
	fmt.Println("Person created: ", p1)
	p1.changeMe()
	fmt.Println("Person changed: ", p1)
}