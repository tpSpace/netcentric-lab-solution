package main

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	length := 10
	var a []byte
	var b []byte

	for i:= 0; i < 10000; i++ {
		a, b = GenerateDNA(length)
		result, err := Distance(a, b)
		if err != nil {
			t.Errorf("Got error")
		} else {
			t.Log(result)
			fmt.Printf("Strands A: %s\n", a)
			fmt.Printf("Strands b: %s\n", b)
			fmt.Printf("Result: %d\n\n", result)
		}
	}
}