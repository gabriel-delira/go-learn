package main

import (
	"fmt"
)

func main()  {
	Test1()
	Test2()
	Test3()
	Test4()
}

type person struct {
	firstName string
	lastName string
	favoritesIcecream []string
}

func Test1(){
	p1 := person{
		firstName: "Seth",
		lastName: "Dragunov",
		favoritesIcecream: []string{"Mint", "Chocolate"},
	}
	p2 := person{
		firstName: "Doc",
		lastName: "Melie",
		favoritesIcecream: []string{"Strawberry", "Cream"},
	}
	fmt.Println("First person: ", p1)
	fmt.Println("Second person: ", p2)

	fmt.Println(p1.firstName, p1.lastName)
	for _, fi := range p1.favoritesIcecream {
		fmt.Println("\t", fi)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for _, fi := range p2.favoritesIcecream {
		fmt.Println("\t", fi)
	}

}

func Test2(){
	p1 := person{
		firstName: "Seth",
		lastName: "Dragunov",
		favoritesIcecream: []string{"Mint", "Chocolate"},
	}
	p2 := person{
		firstName: "Doc",
		lastName: "Melie",
		favoritesIcecream: []string{"Strawberry", "Cream"},
	}

	personsMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, someone := range personsMap {
		fmt.Println(someone.firstName, someone.lastName)
		for _, fi := range someone.favoritesIcecream {
			fmt.Println("\t", fi)
		}
	}

}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func Test3(){
	myTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println("Truck: ", myTruck)
	fmt.Println("Sedan: ", mySedan)

	fmt.Println(myTruck.doors, myTruck.fourWheel, myTruck.vehicle.color)
	fmt.Println(mySedan.color, mySedan.luxury, mySedan.vehicle.color)

}

func Test4(){
	anony := struct{
		someMap map[string]int
		someList []int
		name string
	}{
		name: "Some struct",
		someMap: map[string]int{
			"Way of kings": 8,
			"Name of the wind": 10,
		},
		someList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	fmt.Println(anony)
	fmt.Println(anony.name)
	fmt.Println(anony.someMap)
	fmt.Println(anony.someList)
}