package main

import (
	"fmt"
	"testing"
)

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func ExampleSalutations() {
	fmt.Println("hello, and goodbye")
	// Output:
	// hello, and goodbye
}
