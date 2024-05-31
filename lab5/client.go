package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:12345")
	conn, _ := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Registration
	conn.Write([]byte("JOIN " + username))

	go func() {
		// Receive messages
		buffer := make([]byte, 1024)
		for {
			n, _, _ := conn.ReadFromUDP(buffer)
			fmt.Println(string(buffer[:n]))
		}
	}()

	// Send messages
	for {
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
	}
}
