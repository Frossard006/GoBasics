package sum

import (
	quote "Seventh/foo"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	x := UseCount("eu e você, você você")
	for i, v := range x {
		switch i {
		case "eu":
			if v != 1 {
				t.Error("Got:", v, "Expected:", 1)
			}
		case "e":
			if v != 1 {
				t.Error("Got:", v, "Expected:", 1)
			}
		case "você,":
			if v != 1 {
				t.Error("Got:", v, "Expected:", 1)
			}
		case "você":
			if v != 2 {
				t.Error("Got:", v, "Expected:", 2)
			}
		}
	}
}

func TestCount(t *testing.T) {
	x := "eu e você, amor"
	if Count(x) != 4 {
		t.Error("Got:", x, "Expected:", 4)
	}
}

func ExampleCount() {
	fmt.Println(Count("eu e você, amor"))
	// Output:
	// 4
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
