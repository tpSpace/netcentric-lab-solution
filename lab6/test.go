package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)
func main() {
	// Listen on TCP port 9999
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is listening on port 9999")
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	request, err := reader.ReadString('\n')
	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received request: ", request)

	requestLine := strings.Fields(request)

	if len(requestLine) < 2 {
		fmt.Println("Invalid request")
		return
	}
	url := requestLine[1]

	if url == "/index.html" {
		SendHTMLResponse(conn, "index.html")
	} else if strings.HasSuffix(url, ".jpg") {
		SendFileResponse(conn, url[1:], "image/jpeg")
	} else if strings.HasSuffix(url, ".mp3") {
		SendFileResponse(conn, url[1:], "audio/mpeg")
	} else {
		SendNotFoundResponse(conn)
	}
}

func SendHTMLResponse(conn net.Conn, filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		SendNotFoundResponse(conn)
		return
	}
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: " + strconv.Itoa(len(content)) + "\r\n" +
		"Connection: close\r\n" +
		"\r\n" + string(content)
	conn.Write([]byte(response))
}

func SendFileResponse(conn net.Conn, filename string, contentType string) {
	content, err := os.ReadFile(filename)
	fmt.Print(filename)
	if err != nil {
		SendNotFoundResponse(conn)
		return
	}

	// encode the image to base64
	// encoded := base64.StdEncoding.EncodeToString(content)
	
	// use dynamic date
	response := "HTTP/1.1 200\r\n" +
		"Content-Type: " + contentType + "\r\n" +
		"Content-Length: " + strconv.Itoa(len(content)) + "\r\n" +
		"Connection: close\r\n" +
		"\r\n"
	conn.Write([]byte(response))
	conn.Write(content)
}

func SendNotFoundResponse(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		"404 Not Found"
	conn.Write([]byte(response))
}
