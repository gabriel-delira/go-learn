package main

import (
	"fmt"
	"math"
)

func main()  {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
	Test7()
	Test8()
	Test9()
	Test10()
}

func foo1() int {
	return 2
}

func bar1() (string, int) {
	return "test", 10
}

func Test1(){
	x := foo1()
	y, z := bar1()
	fmt.Println("foo =>", x)
	fmt.Printf("bar => %s AND %d\n", y, z)
}

func foo2(xi ...int) int {
	total := 0
	for _, v := range xi{
		total += v
	}
	return total
}

func bar2(xi []int) int {
	total := 0
	for _, v := range xi{
		total += v
	}
	return total
}

func Test2(){
	fmt.Println("Sum:", foo2([]int{1, 2, 3, 4, 5}...))
	fmt.Println("Sum 2:", bar2([]int{1, 2, 3, 4, 5}))
}

func foo3()  {
	defer func(){fmt.Println("============ END 3 ============")}()
	fmt.Println("========== MIDDLE 3 ==========")
}

func Test3(){
	defer foo3()
	fmt.Println("========= Begin test3 =========")
}

type person struct {
	first string
	last string
	age int
}

func (p person)speak(){
	fmt.Printf("My name is %s %s and I am %d years old. :)\n", p.first, p.last, p.age)
}

func Test4(){
	p1 := person{
		first: "Ruby",
		last: "Stark",
		age: 17,
	}
	p1.speak()
}

type square struct{
	side float64
}
type circle struct{
	radius float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}
func (c circle) area() float64 {
	return math.Pi*math.Pow(c.radius, 2)
}

type shape interface{
	area() float64
}

func info(sh shape)  {
	fmt.Println("Area", sh.area())
}

func Test5(){
	s := square{ side: 4 }
	c := circle{ radius: 2 }

	info(s)
	info(c)
}

func Test6(){
	func(){
		fmt.Println("I am an anonymous function")
	}()
}

var t7 func()

func Test7(){
	anonyF := func(){
		fmt.Println("I am an anonymous function Assigned to a variable")
	}
	t7 = anonyF
	fmt.Printf("%T\n", anonyF)
	anonyF()
	t7()
}

func T8F() func() {
	return func(){
		fmt.Println("I am a func inside a func how cool is that ?!!!")
	}
}

func T8F2() func() int {
	return func() int {
		return 2 + 3
	}
}

func Test8(){
	f8 := T8F()
	f8()

	f82 := T8F2()
	value := f82()
	fmt.Println("A value returned by a func insisde a func IS", value)
}

func T9F(f func()) {
	f()
}
func T9F2(f func() int) int {
	return f()
}
func T9F3(f func(v int)) {
	f(5)
}
func T9F4(f func(v int)string) string {
	return f(7)
}

func Test9(){
	f91 := func(){ fmt.Println(" Begin of test 9, func as parameters") }
	f92 := func() int { return 10 }
	f93 := func(v int){ fmt.Printf(" Received param plus 20: %d\n", v+20) }
	f94 := func(v int) string { return fmt.Sprintf(" Received param: %v\n", v) }

	T9F(f91)
	fr92 := T9F2(f92)
	T9F3(f93)
	fr94 := T9F4(f94)
	fmt.Println("T9F2", fr92)
	fmt.Println("T9F4", fr94)

}

func foo10() func() int {
	x := 10
	return func() int {
		x *= 2
		return x
	}
}

func Test10(){
	incrementor := foo10()
	fmt.Println("x=", incrementor())
	fmt.Println("x=", incrementor())
	fmt.Println("x=", incrementor())
	fmt.Println("x=", incrementor())
	fmt.Println("x=", incrementor())
	incrementor2 := foo10()
	fmt.Println("y=", incrementor2())
	fmt.Println("y=", incrementor2())
	fmt.Println("y=", incrementor2())
	fmt.Println("y=", incrementor2())
	fmt.Println("y=", incrementor2())
	fmt.Println("y=", incrementor2())
}