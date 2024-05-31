package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	
	login := false
	// Resolve TCP address
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	//
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	for {
		if login {
			// buffer to get data
			received := make([]byte, 1024)
			str := string(received)
			_, err = conn.Read(received)

			// error 
			if err != nil && err.Error() != "EOF" {
				println("Read data failed:", err.Error())
				os.Exit(1)
			}		

			// print message
			if str != ""  {
				fmt.Print("Received message: " + str + "\n")
			}
			if str == "start" {
				for {
					// start the geussing game
					fmt.Print("Guess the number: ")
					var guess int
					fmt.Scan(&guess)
					_, err = conn.Write([]byte(fmt.Sprintf("%d", guess)))
					if err != nil {
						println("Write data failed:", err.Error())
						os.Exit(1)
					}
					// buffer to get data
					received := make([]byte, 1024)
					_, err = conn.Read(received)
					if err != nil {
						println("Read data failed:", err.Error())
						os.Exit(1)
					}
					// print message
					fmt.Print("Received message: " + string(received) + "\n")
					if string(received) == "correct" {
						break
					}
				}
			}
			// clear buffer
			str = ""
			received = make([]byte, 1024)

		} else {
			fmt.Print("1. Login\n2. Register\n3. Exit\n")
			var choice int
			fmt.Scan(&choice)
			// reader := bufio.NewReader(os.Stdin)
			switch choice {
			case 1:
				fmt.Print("Username: ")
				var username string
				// read the input until "Enter" is pressed exclude the newline character
				fmt.Scan(&username)
				
				fmt.Print("Password: ")
				var password string
				// read the input until "Enter" is pressed exclude the newline character
				fmt.Scan(&password)

				_, err = conn.Write([]byte("login:" + username + ":" + password+ ":"))
				if err != nil {
					println("Write data failed:", err.Error())
					fmt.Printf("here\n")
					os.Exit(1)
				}

				// read response
				received := make([]byte, 1024)
				_, err = conn.Read(received)
				if err != nil {
					println("Read data failed:", err.Error())
					os.Exit(1)
				}
				fmt.Printf("Received message: %s\n", string(received))
				fmt.Println("Received message: ", len(strings.TrimSpace(string(received))))
				if strings.Compare(string(received), "login success") == 0{
					login = true
				} else {
					login = false
				}
			case 2:
				fmt.Print("Username: ")
				var username string
				fmt.Scan(&username)
				fmt.Print("Password: ")
				var password string
				fmt.Scan(&password)
				_, err = conn.Write([]byte("register:" + username + ":" + password))
				if err != nil {
					println("Write data failed:", err.Error())
					os.Exit(1)
				}
			case 3:
				os.Exit(1)
			}
			
		}
	}
	
}
