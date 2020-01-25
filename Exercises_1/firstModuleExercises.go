package main

import (
	"fmt"
	"reflect"
)

var x2 int;
var y2 string;
var z2 bool;

func main()  {
	Test1()
	Test2()
	Test3()
	Test2()

	Test4()
	Test5()
	Test6()
}

func Test1(){
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(" EXERCISE 1 ")
	fmt.Println("")


	fmt.Println("First test: ")
	fmt.Println(x, y, z)
	fmt.Println("Another way")
	fmt.Printf("x = %d \ny = %s \nz = %t \n \n", x, y, z)

	fmt.Println("Second test: ")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("Another way")
	fmt.Printf("x = %d \n", x)
	fmt.Printf("y = %s \n", y)
	fmt.Printf("z = %t \n", z)

	fmt.Println("End of exercise 1")
	fmt.Println("  ========================================  ")
}

func Test2()  {
	fmt.Println(" EXERCISE 2 ")
	fmt.Println("")
	fmt.Printf("x2 = %d \ny2 = %s \nz2 = %t \n \n", x2, y2, z2)	
	fmt.Println("End of exercise 2")
	fmt.Println("  ========================================  ")
}

func Test3(){
	x2 = 42
	y2 = "James Bond"
	z2 = true

	s := fmt.Sprintf("Now the values are:\nx2 = %v, y2 = %v, z2 = %v ", x2, y2, z2)

	fmt.Println(" EXERCISE 3 ")
	fmt.Println(s)
	fmt.Println("End of exercise 3")
	fmt.Println("  ========================================  ")
}

type MyFirstModuleType int
var x4 MyFirstModuleType

func Test4(){

	fmt.Println(" EXERCISE 4 ")
	fmt.Printf("Current value of x4: %v\n", x4)
	fmt.Printf("[using reflect] Type of x4: %v\n", reflect.TypeOf(x4))
	fmt.Printf("[using Printf] Type of x4: %T\n", x4)

	x4 = 42

	fmt.Printf("New value of x4: %v\n", x4)

	fmt.Println("End of exercise 4")
	fmt.Println("  ========================================  ")
}

var y5 int

func Test5()  {
	y5 = int(x4)
	
	fmt.Println(" EXERCISE 5 ")
	fmt.Printf("Current value of y5: %v\n", y5)
	fmt.Printf("Type of y5: %T\n", y5)

	fmt.Println("End of exercise 5")
	fmt.Println("  ========================================  ")

}


func Test6() {
	fmt.Println(" EXERCISE 6 ")

	fmt.Println("End of exercise 6")
	fmt.Println("  ========================================  ")
}