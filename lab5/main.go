package main

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Name string
	Addr *net.UDPAddr
}

var clients = make(map[string]*Client)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":12345")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, _ := conn.ReadFromUDP(buffer)
		message := string(buffer[:n])

		// Client registration
		if strings.HasPrefix(message, "JOIN ") {
			name := strings.TrimPrefix(message, "JOIN ")
			clients[name] = &Client{name, clientAddr}
			fmt.Println("Client joined:", name)
			continue
		}

		// Message handling
		parts := strings.SplitN(message, " ", 2)
		if len(parts) == 2 {
			target := parts[0]
			text := parts[1]

			// Broadcast
			if target == "@all" {
				for _, c := range clients {
					conn.WriteToUDP([]byte(text), c.Addr)
				}
			} else if strings.HasPrefix(target, "@") {
				// Private message
				targetName := strings.TrimPrefix(target, "@")
				if client, ok := clients[targetName]; ok {
					conn.WriteToUDP([]byte(text), client.Addr)
				}
			}
		}
	}
}
