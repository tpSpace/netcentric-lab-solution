// Package main provides a simple HTTP server.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// main starts the server and listens for connections on TCP port 9999.
func main() {
    // Listen on TCP port 9999
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer ln.Close()
    fmt.Println("Server is listening on port 9999")

    // Accept connections in a loop and handle each one in a separate goroutine.
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        go handleConnection(conn)
    }
}

// handleConnection handles an individual client connection.
// It reads the request, parses it, and sends the appropriate response.
func handleConnection(conn net.Conn) {
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
        sendNotFoundResponse(conn)
        return
    }
    method := requestLine[0]
    url := requestLine[1]

    if method != "GET" {
        fmt.Println("Unsupported method")
        sendNotFoundResponse(conn)
        return
    }

    // Send the appropriate response based on the requested URL.
    if url == "/index.html" {
        sendHTMLResponse(conn, "index.html")
    } else if strings.HasSuffix(url, ".jpg") {
        sendFileResponse(conn, url[1:], "image/jpeg")
    } else if strings.HasSuffix(url, ".mp3") {
        sendFileResponse(conn, url[1:], "audio/mpeg")
    } else if strings.HasSuffix(url, ".ico") {
        sendFileResponse(conn, url[1:], "image/x-icon")	
    } else {
        sendNotFoundResponse(conn)
    }
}

// sendHTMLResponse sends an HTML response to the client.
// It reads the specified file and sends it as the response body.
func sendHTMLResponse(conn net.Conn, filename string) {
    content, err := os.ReadFile(filename)
    if err != nil {
        sendNotFoundResponse(conn)
        return
    }
    response := "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/html\r\n" +
        "Content-Length: " + strconv.Itoa(len(content)) + "\r\n" +
        "Connection: close\r\n" +
        "\r\n" + string(content)
    conn.Write([]byte(response))
}

// sendFileResponse sends a file response to the client.
// It reads the specified file and sends it as the response body.
func sendFileResponse(conn net.Conn, filename string, contentType string) {
    content, err := os.ReadFile(filename)
    if err != nil {
        sendNotFoundResponse(conn)
        return
    }

    response := "HTTP/1.1 200 OK\r\n" +
        "Content-Type: " + contentType + "\r\n" +
        "Content-Length: " + strconv.Itoa(len(content)) + "\r\n" +
        "Connection: close\r\n" +
        "\r\n"
    conn.Write([]byte(response))
    conn.Write(content)
}

// sendNotFoundResponse sends a 404 Not Found response to the client.
func sendNotFoundResponse(conn net.Conn) {
    response := "HTTP/1.1 404 Not Found\r\n" +
        "Content-Type: text/plain\r\n" +
        "Content-Length: 13\r\n" +
        "Connection: close\r\n" +
        "\r\n" +
        "404 Not Found"
    conn.Write([]byte(response))
}