package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"
	"math/rand"
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
	fmt.Println("Begin - cpus:", runtime.NumCPU())
	fmt.Println("Begin - go routines:", runtime.NumGoroutine())

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func ()  {
		fmt.Println("My First goroutine print")
		wg.Done()
	}()
	go func ()  {
		fmt.Println("My second goroutine print")
		wg.Done()
	}()
	
	fmt.Println("Middle - go routines:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End - cpus:", runtime.NumCPU())
	fmt.Println("End - go routines:", runtime.NumGoroutine())
}

type Person struct {
	FirstName string
	LastName string
	Speaks []string
}

func (p *Person) Speak(speakIndex int) {
	if speakIndex < 0 {
		speakIndex = rand.Intn(len(p.Speaks)) // Get a random phrase from speaks list
	}
	fmt.Println(p.Speaks[speakIndex]) 
}

type Human interface {
	Speak(speakIndex int)
}

func SaySomething(h Human) {
	h.Speak(0)
}

func Test2(){
	p1 := Person{
		FirstName: "John",
		LastName: "Summers",
		Speaks: []string{"My name is John Summers", },
	}

	SaySomething(&p1)
	//SaySomething(p1) // This cannot be used because it is not a pointer and the Method Speak expects a pointer
}

func Test3(){
	// To show a race condition : go run -race main.go
	// To show a race condition on builded program : go build -race main.go THEN execute generated file in our case ./main
	counter := 0
	numberOfGoroutines := 100
	wg := sync.WaitGroup{}
	wg.Add(numberOfGoroutines)

	for i := 1; i <= numberOfGoroutines; i++ {
		go func () {
			j := counter
			runtime.Gosched()
			j++
			counter = j
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter is", counter)
}

func Test4(){
	counter := 0
	numberOfGoroutines := 100
	wg := sync.WaitGroup{}
	wg.Add(numberOfGoroutines)
	mux := sync.Mutex{}

	for i := 1; i <= numberOfGoroutines; i++ {
		go func () {
			mux.Lock()
			j := counter
			// runtime.Gosched() // This line can be removed because doesnt make sense in this case
			j++
			counter = j
			mux.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter is", counter)
}

func Test5(){
	// Using atomic
	var counter int64 = 0
	numberOfGoroutines := 100
	wg := sync.WaitGroup{}
	wg.Add(numberOfGoroutines)	

	for i := 1; i <= numberOfGoroutines; i++ {
		go func () {						
			// runtime.Gosched() // This line can be removed because doesnt make sense in this case
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter now is :", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter is", counter)
}

func Test6(){
	// go run <programFileName> => run the program
	// go build <programFileName> => build the program and generate the built file. After that is just execute the program using its path
	// eg: program is "Abacaxi" and folder location is /home/my/path then to run built program is './home/my/path/Abacaxi'
	// go install <programFileName> => install the program into "bin" folder of "go" folder [GOBIN]. After that is just call the program by its name
	// eg: program is "Abacaxi" then to run installed program is 'Abacaxi'
	fmt.Printf("OS: %s\nArchiteture: %s\n", runtime.GOOS, runtime.GOARCH)
}

func Test7(){
	// ...........
}