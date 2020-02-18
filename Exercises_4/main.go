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
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
func Test2() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slice)
}
func Test3() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(slice[:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])
	fmt.Println(slice[1:6])
}
func Test4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
		
}
func Test5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[0:3], x[6:]...)
	fmt.Println(x)
}
func Test6() {
	states := make([]string, 27, 27)
	states = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	fmt.Printf("Length: %v\n", len(states))
	fmt.Printf("Capacity: %v\n", cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Printf("Index: %d, Value: %s \n", i, states[i])
	}
}
func Test7() {
	x := [][]string{}
	y := []string{ "Merry", "Go", "Sunny", "Natsu" }
	x = append(x, []string{ "James", "Bond", "t1", "t2" })
	x = append(x, y)
	for _, v := range x {
		fmt.Println("Record: ", v)
		for i, w := range v {
			fmt.Printf("Index: %v, Value: %s\n", i, w)
		}
	}
}
func Test8() {
	m := map[string][]string{
		"ln1": []string{ "books", "films", "dance" },
		"ln2": []string{ "cycling", "basket", "music" },
		"ln3": []string{ "tenis", "animes", "cars" },
	}
	for k, vs := range m {
		fmt.Println("Key: ", k)
		for i, v := range vs {
			fmt.Printf("\tIndex: %d, Value: %s\n", i, v)
		}
	}
}
func Test9() {
	m := map[string][]string{
		"ln1": []string{ "books", "films", "dance" },
		"ln2": []string{ "cycling", "basket", "music" },
		"ln3": []string{ "tenis", "animes", "cars" },
	}

	m["New_ln"] = []string{ "books", "animes", "games" }

	for k, vs := range m {
		fmt.Println("Key: ", k)
		for i, v := range vs {
			fmt.Printf("\tIndex: %d, Value: %s\n", i, v)
		}
	}
}
func Test10() {
	m := map[string][]string{
		"ln1": []string{ "books", "films", "dance" },
		"ln2": []string{ "cycling", "basket", "music" },
		"ln3": []string{ "tenis", "animes", "cars" },
	}

	delete(m, "ln2")

	for k, vs := range m {
		fmt.Println("Key: ", k)
		for i, v := range vs {
			fmt.Printf("\tIndex: %d, Value: %s\n", i, v)
		}
	}
}