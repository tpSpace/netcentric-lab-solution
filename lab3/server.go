package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Fullname string   `json:"fullname"`
	Email    []string `json:"email"`
	Address  []string `json:"address"`
}

func main() {

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// print message
	fmt.Println("Received message:", string(buffer))

	mess := string(buffer)

	if mess[:4] == "exit" {
		fmt.Println("exit")
		_, err = conn.Write([]byte("exit success"))
		os.Exit(1)
	} else if mess[:5] == "login" {
		
		// read user.json and check if user exists
		jsonFile, err := os.Open("user.json")

		// check if file exists
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := io.ReadAll(jsonFile)
		var users []User

		json.Unmarshal(byteValue, &users)

		// extract using plsit ":"
		data := strings.Split(mess, ":")

		for _, user := range users {
			username := data[1]
			password := data[2]

			// check if username and password match
			if user.Username == username && user.Password == password {
				fmt.Println("login success")
				_, err = conn.Write([]byte("login success\n"))
				// delay wait until the message is sent
				time.Sleep(1 * time.Second)
				// _, err = conn.Write([]byte("Welcome " + user.Fullname + "\n"))
					if err != nil {
						log.Fatal(err)
					}
				// send message
				_, err = conn.Write([]byte("Let's start the game!\n"))
				if err != nil {
					log.Fatal(err)
				}
				// send start message
				_, err = conn.Write([]byte("start"))
				if err != nil {
					log.Fatal(err)
				}
				// random a number
				rand.Seed(time.Now().UnixNano())
				number := rand.Intn(100)
				for {
					// send message
					_, err = conn.Write([]byte("Enter your guess: "))
					if err != nil {
						log.Fatal(err)
					}
					// buffer to get data
					received := make([]byte, 1024)
					_, err = conn.Read(received)
					if err != nil {
						log.Fatal(err)
					}
					// convert to string
					str := string(received)
					// convert to int
					guess := 0
					fmt.Sscanf(str, "%d", &guess)
					// check if guess is correct
					if guess == number {
						_, err = conn.Write([]byte("Congratulations! You got it! "))
						if err != nil {
							log.Fatal(err)
						}
						break
					} else if guess < number {
						_, err = conn.Write([]byte("Too low! "))
						if err != nil {
							log.Fatal(err)
						}
					} else {
						_, err = conn.Write([]byte("Too high! "))
						if err != nil {
							log.Fatal(err)
						}
					}
					
				}
			} 
		}
		


		if err != nil {
			log.Fatal(err)
		}
	} else if mess[:8] == "register" {
		fmt.Println("register")
		_, err = conn.Write([]byte("register success"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("invalid")
		_, err = conn.Write([]byte("invalid"))
		if err != nil {
			log.Fatal(err)
		}
	}

	// close conn
	conn.Close()
}
