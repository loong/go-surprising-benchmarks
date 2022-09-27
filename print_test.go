package main

import (
	"os"
	"fmt"
	"testing"
)

// fmt.Println("Hello " + worldString)
// vs
// fmt.Println("Hello", worldString)
//
// Result:
// =======
// goos: darwin
// goarch: amd64
// pkg: github.com/loong/go-surprising-benchmarks
// cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
// BenchmarkConcat-8       95025295               111.5 ns/op
// BenchmarkNativeFmt-8    151688880               79.84 ns/op
// PASS
// ok      github.com/loong/go-surprising-benchmarks       31.166s

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
