package main

import (
	"bufio"
	"fmt"
	"goldis/commands"
	"goldis/resp"
	"net"
	"os"
)

// Goldis server constants
const (
	GoldisHost = "localhost"
	GoldisPort = "6382"
	connType   = "tcp"
)

func main() {

	// Listen for incoming connections.
	l, err := net.Listen(connType, GoldisHost+":"+GoldisPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Server listening at :", GoldisHost, GoldisPort)
	// Close the listener when the application closes.
	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}

}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Create a new reader
	fmt.Println("Client connected")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		inputString, err := resp.ParseRequest(bytes)
		if err == nil {
			result := commands.ExecuteCommand(inputString)
			conn.Write([]byte(result + "\n"))
		}
	}
}
