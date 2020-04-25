package main

import (
	"fmt"
	"sync"
)

func main()  {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
	Test7()
}

func Test1(){
	// First solution:
	// c := make(chan int)
	// ORIGINAL LINE: c <- 42

	// go func (c chan int){
	// 	c <- 42
	// }(c)

	// Second solution:
	c := make(chan int, 1)
	c <- 42
		

	fmt.Println(<-c)
}

func Test2(){
	//Line with problem: cs := make(chan<- int)
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<- cs)

	fmt.Printf("-------\n")
	fmt.Printf("cs\t%T\n", cs)

	//Line with problem: cs2 := make(<-chan int)
	cs2 := make(chan int)

	go func() {
		cs2 <- 42
	}()

	fmt.Println(<- cs2)

	fmt.Printf("-------\n")
	fmt.Printf("cs2\t%T\n", cs2)
}

func gen() <-chan int {
	c := make(chan int)
	
	go func (c chan int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}(c)

	return c
}

func receive(c <-chan int)  {
	for v := range c {
		fmt.Println("Current Value: ", v)
	}
}

func Test3(){
	c := gen()
	receive(c)

	fmt.Println("about to exit Test 3")
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)
	
	go func () {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive2(c <-chan int, q <-chan int)  {
	for {
		select {
		case v := <-c:
			fmt.Println("Even", v)
		case <-q:
			fmt.Println("Finish")
			return
		}
	}
}

func Test4(){
	fmt.Println("Test 4")
	q := make(chan int)
	c := gen2(q)

	receive2(c, q)

	fmt.Println("about to exit Test 4")
}

func Test5(){
	c := make(chan int, 1)
	c <- 42

	// ANOTHER SOLUTION
	//c := make(chan int)
	// go func(){c <- 42}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func Test6(){
	c := make(chan int)

	go func ()  {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func Test7(){
	fmt.Println("Test 7")
	c := make(chan int)
	numberOfRoutines := 10
	wg := sync.WaitGroup{}
	wg.Add(numberOfRoutines)

	func (){
		for i := 0; i < numberOfRoutines; i++ {
			go func (iterator int) {
				for j := 0; j < 10; j++ {
					c <- j+(iterator*10)
				}
				wg.Done()
			}(i)			
		}
		wg.Wait()
		close(c)
	}()	

	for v := range c {
		fmt.Println(v)
	}
}