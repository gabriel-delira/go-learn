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
}

func Test1()  {
	fmt.Println(" ===================== ")
	fmt.Println(" Ex. 1 ")

	x := 64
	fmt.Printf(" Binary: %b\n", x)
	fmt.Printf(" Decimal: %d\n", x)
	fmt.Printf(" Hex: %X\n", x)
	fmt.Printf(" Hex 2: %#x\n", x)
	fmt.Println()
}

func Test2()  {
	fmt.Println(" ===================== ")
	fmt.Println(" Ex. 2 ")

	x := 64 == 30
	y := 64 <= 30
	z := (64 >= 30)
	w := 64 != 30
	r := (64 < 30)
	s := (64 > 30)

	fmt.Printf(" x := 64 == 30   => %v\n", x)
	fmt.Printf(" y := 64 <= 30   => %v\n", y)
	fmt.Printf(" z := (64 >= 30) => %v\n", z)
	fmt.Printf(" w := 64 != 30   => %v\n", w)
	fmt.Printf(" r := (64 < 30)  => %v\n", r)
	fmt.Printf(" s := (64 > 30)  => %v\n", s)
	
	fmt.Println()
}

const (
	x3 = 30
	x3l int = 30
)

func Test3()  {
	fmt.Println(" ===================== ")
	fmt.Println(" Ex. 3 ")
	fmt.Printf(" Untyped const x3: %v\n", x3)
	fmt.Printf(" Typed const x3l: %v\n", x3l)
	fmt.Println()
}

func Test4()  {
	fmt.Println(" ===================== ")
	fmt.Println(" Ex. 4 ")

	x := 64
	fmt.Printf(" Binary: %b\n", x)
	fmt.Printf(" Decimal: %d\n", x)	
	fmt.Printf(" Hex: %#x\n", x)

	fmt.Println("After shit")
	y := 64 << 1
	fmt.Printf(" Binary: %b\n", y)
	fmt.Printf(" Decimal: %d\n", y)	
	fmt.Printf(" Hex: %#x\n", y)

	fmt.Println()
}

func Test5()  {
	fmt.Println(" ===================== ")
	fmt.Println(" Ex. 5 ")

	x := `This is 
	a multiline
	string variable 
	"Cool isn't it?"`
	
	fmt.Println(x)
	fmt.Println()
}

const (
	fourYearsAgo = 2016 + iota
	threeYearsAgo = 2016 + iota
	twoYearsAgo = 2016 + iota
	oneYearsAgo = 2016 + iota
)

func Test6()  {
	fmt.Println(" ===================== ")
	fmt.Println(" Ex. 6 ")
	fmt.Printf(" (2015 + iota) Year is  %v\n", fourYearsAgo)
	fmt.Printf(" (2015 + iota) Year is  %v\n", threeYearsAgo)
	fmt.Printf(" (2015 + iota) Year is  %v\n", twoYearsAgo)
	fmt.Printf(" (2015 + iota) Year is  %v\n", oneYearsAgo)
	fmt.Println()
}