package word

import(
	"testing"
	"fmt"
)

func ExampleCount()  {
	fmt.Println(Count("test"))
	// Output
	// 1	
}


func TestUseCount(t *testing.T)  {
	vMap := UseCount("test")
	if vMap["test"] != 1 {
		t.Error("Wrong count")
	}
}

func TestCount(t *testing.T)  {
	v := Count("test")
	if v != 1 {
		t.Error("Wrong count")
	}
}

func BenchmarkUseCount(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		UseCount("test")
	}
}

func BenchmarkCount(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Count("test")
	}
}