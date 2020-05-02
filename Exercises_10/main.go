package main

import (
	"fmt"
	"encoding/json"
	"log"
	"errors"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}


func main()  {
	Test1()
	Test2()
	Test3()
	Test4()
}

func Test1(){
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, errMarshal := json.Marshal(p1)
	if errMarshal != nil {
		log.Fatalln("Error:", errMarshal)
	}
	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	return bs, err
}

func Test2(){
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error during json readind: %v", err))
	}

	fmt.Println(string(bs))
}

type customError struct {
	Message string
	Err error 
}

func (ce customError) Error() string {
	return fmt.Sprintf("%v:\n\t%v", ce.Message, ce.Err)
}

func foo(err error) customError {
	fmt.Println(err.Error())
	return customError{Message: "Some error has occured, please check below:", Err: err}
}

func Test3(){
	e1 := customError{Message: "error test 3", Err: errors.New("WTF")}
	foo(e1)	
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := sqrtError{
			lat: "50.2289 N",
			long: "99.4656 W",
			err: errors.New(fmt.Sprintf("Cannot do sqrt of negative number [%v], please use another value.", f)),
		}
		return 0, err
	}
	return 42, nil
}

func Test4(){
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}
