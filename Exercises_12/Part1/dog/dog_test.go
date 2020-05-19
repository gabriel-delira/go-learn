package dog

import(
	"testing"
	"fmt"
)

func ExampleYears()  {	
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo()  {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}


func TestYears(t *testing.T)  {	
	y := Years(10)
	if y != 70 {
		t.Error("Wrong number")	
	}
	
}

func TestYearsTwo(t *testing.T)  {
	y := YearsTwo(10)
	if y != 70 {
		t.Error("Wrong number")
	}
}


func BenchmarkYears(b *testing.B)  {	
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B)  {	
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}