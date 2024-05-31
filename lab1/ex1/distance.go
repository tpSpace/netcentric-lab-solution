package main

import (
	"errors"
	"math/rand"
)

func Distance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("STRANDS SHOULD HAVE SAME LENGTH")
	}

	distance := 0;
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	} 

	return distance, nil;
}

func GenerateDNA(length int) ([]byte, []byte) {
	const dna = "ATGC"
	a := make([]byte, length)
	b := make([]byte, length)
	for i := range a {
		a[i] = dna[rand.Intn(len(dna))]
		b[i] = dna[rand.Intn(len(dna))]
	}
	return a, b
}

