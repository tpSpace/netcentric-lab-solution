package main

import "fmt"

func main() {
	
	a := "lfdkajlkfjwekflhe"
	sum := 0
	fmt.Printf("Your input: %s\n", a);
	for _, c := range a {
		sum += score(string(c))
	}
	fmt.Printf("Your sum is: %d\n", sum)
}

func score(letter string) (int) {
	switch letter {
		case "d", "g" : return 2
		case "b", "c", "m", "p" : return 3
		case "f", "h", "v", "w", "y": return 4
		case "k" : return 5
		case "j", "x" : return 8
		case "q", "z" : return 10
	default: return 1
	}
}