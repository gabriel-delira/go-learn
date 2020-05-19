package mymath

import(
	"testing"
	"fmt"
)

func ExampleCenteredAvg()  {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000}
	fmt.Println(CenteredAvg(input))
	// Output
	// 5.5
}

func TestCenteredAvg(t *testing.T)  {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000}
	v := CenteredAvg(input)
	if v != 5.5 {
		t.Error("Wrong count")
	}
}


func BenchmarkCenteredAvg(b *testing.B)  {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000}
	for i := 0; i < b.N; i++ {
		CenteredAvg(input)
	}
}