package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func isValidNumber(input string) bool {
    input = strings.ReplaceAll(input, " ", "")

    sum := 0
    doubleNext := false

    for i := len(input) - 1; i >= 0; i-- {
        digit, _ := strconv.Atoi(string(input[i]))

        if doubleNext {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
        doubleNext = !doubleNext 
    }
    return sum%10 == 0
}

func GenerateNumber(length int) (string) {
    const letters = "0123456789"
    a := make([]byte, length)
    for i := range a {
        a[i] = letters[rand.Intn(len(letters))]
    }
    return string(a)
}

func main () {
    for i := 0; i < 100; i++ {
        input := GenerateNumber(16)
        fmt.Printf("Attempt number: %b\n", i)
        if isValidNumber(input) {
            fmt.Printf("The number %s is valid!\n", input)
        } else {
            fmt.Printf("The number %s is not valid.\n", input)
        }
    }
}