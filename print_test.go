package main

import (
	"os"
	"fmt"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	temp := os.Stdout
	os.Stdout = nil    // turn it off
	
	worldString := "World"
	for i := 0; i < b.N; i++ {
		fmt.Println("Hello " + worldString)
	}

	os.Stdout = temp   // restore it
}	

func BenchmarkNativeFmt(b *testing.B) {
	temp := os.Stdout
	os.Stdout = nil    // turn it off

	worldString := "World"
	for i := 0; i < b.N; i++ {
		fmt.Println("Hello", worldString)
	}

	os.Stdout = temp   // restore it
}
