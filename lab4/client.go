package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Println("Connected to Hangman Game Server")

	// Authentication
	username := prompt("Enter username: ")
	password := prompt("Enter password: ")

	authData := fmt.Sprintf("%s_%s\n", username, password)
	conn.Write([]byte(authData))

	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if strings.TrimSpace(response) != "authenticated" {
		fmt.Println("Authentication failed.")
		return
	}
	fmt.Println("Authentication successful. Starting game.")

	// Game loop
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(message)

		if strings.Contains(message, "Your turn") {
			guess := prompt("Enter your guess (one letter): ")
			conn.Write([]byte(guess + "\n"))
		}

		if strings.Contains(message, "Game over") {
			break
		}
	}
}

func prompt(promptMsg string) string {
	fmt.Print(promptMsg)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}
