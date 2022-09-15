package fatorial

import (
	"fmt"
	"testing"
)

var value int = 60

func BenchmarkRecursiveFat(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = RecursiveFat(value)
	}
	fmt.Println(r)
}

func BenchmarkIterativeFat(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = IterativeFat(value)
	}
	fmt.Println(r)
}

func BenchmarkConcurrencyFatBySort(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = ConcurrencyFatBySort(value)
	}
	fmt.Println(r)
}

func BenchmarkConcurrencyFat(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = ConcurrencyFat(value)
	}
	fmt.Println(r)
}
