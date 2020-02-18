package main

import (
	"fmt"
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

func Test1(){
	for i := 1; i < 10001; i++ {
		fmt.Println(i)
	}
}
func Test2() {
	for i := 65; i < 91; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
func Test3() {
	i := 1992
	c := 1992
	for c <= 2020 {
		fmt.Println(c-i)
		c++
	}
}
func Test4() {
	i := 1992
	c := 1992
	for {
		fmt.Println(c-i)
		c++
		if c > 2020 {
			break
		}
	}	
}
func Test5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Number: %d \t divided by 4 so remains: %d\n", i, i % 4)
	}
}
func Test6() {
	x := "test"
	if x == "test" {
		fmt.Println("Uhhhhuuuu !!!")
	}
}
func Test7() {
	x := "test" // Just change this value to test all cases
	if x == "test" {
		fmt.Println("Uhhhhuuuu !!! [test]")
	} else if x == "Something" {
		fmt.Println("Doing ...")
	} else {
		fmt.Println("ahhhhhh :(")
	}
}
func Test8() {
	switch {
		case false:
			fmt.Println("FALSE ...")
		case true:
			fmt.Println("TRUE :) ")				
	}
}
func Test9() {
	favSport := "Basket" // Just change this value to test all cases
	switch favSport {
		case "Soccer":
			fmt.Println("Soccer :)")
		case "Football":
			fmt.Println("Football :) ")
		case "Basket":
			fmt.Println("Basket :) ")
		case "Tenis":
			fmt.Println("Tenis :) ")
	}
}
func Test10() {
	fmt.Printf("True && True: %t\n", true && true)
	fmt.Printf("True && False: %t\n", true && false)
	fmt.Printf("True || True: %t\n", true || true)
	fmt.Printf("True || False: %t\n", true || false)
	fmt.Printf("!True: %t\n", !true)
}