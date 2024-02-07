// Server (Windows PC)
// Code to receive Rust code from client, build, and deploy to micro:bit V2

package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

func main() {
	ln, err := net.Listen("tcp", ":port")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle client connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Receive Rust code file from client and save it
	file, err := os.Create("received_code.rs") // Create a new file to save the received Rust code
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, conn) // Copy data from the connection to the file
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return
	}
	fmt.Println("Rust code file received and saved successfully!")

	// Run build script
	cmd := exec.Command("build_script.bat") // Adjust the script name and path as needed
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running build script:", err)
		return
	}

	fmt.Println("Build and deployment successful!")
}
