package main

import (
	"fmt"
	"net"
	"time"
)

const broadcastAddress string = "192.168.1.3:1999"

func BoardcastFileInfo(filename string, filesize int64) {
	for {
		sendBroadcast(filename, filesize)
		time.Sleep(10 * time.Second)
	}
}

func SendBroadcast(filename string, filesize int64) {
	conn, err := net.Dial("udp", broadcastAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := fmt.Sprintf("%s:%d", filename, filesize)

	_, err = conn.Write([]byte(message))

	if err != nil {
		panic(err)
	}
}

func ListenForBroadcast() {
	addr, err := net.ResolveUDPAddr("udp", ":1999")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		message := string(buffer[:n])
		fmt.Println("Received broadcast:", message)
	}
}
