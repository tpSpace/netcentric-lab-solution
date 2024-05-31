package main

import (
	"fmt"
	"strings"
	"sync"
)

func count(word string, result chan<- map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()
	counts := make(map[rune]int)
	for _, r := range word {
		counts[r]++
	}
	result <- counts
}

func main() {
	var wg sync.WaitGroup
	var str string = "we wjwelk weldk wlkedj wledj wel"
	// Add the number of goroutines to wait for

	wordChan := make(chan map[rune]int)
	// split the string by the space
	words := strings.Split(str, " ")
	wg.Add(len(words))
	fmt.Print(len(words), "\n")

	for _, word := range words {
		go count(word, wordChan, &wg)
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(wordChan)
	}()
	<-wordChan

	// combine the results
	result := make(map[rune]int)
	for word := range wordChan {
		for key, value := range word {
			result[key] += value
		}
	}

	// print the result
	for key, value := range result {
		fmt.Printf("%c: %d\n", key, value)
	}
}
