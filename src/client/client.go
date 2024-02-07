package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "server-ip:port")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create a new watcher to monitor directory changes
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		os.Exit(1)
	}
	defer watcher.Close()

	// Add the directory you want to monitor
	dir := "/embedded"
	err = watcher.Add(dir)
	if err != nil {
		fmt.Println("Error adding directory to watcher:", err)
		os.Exit(1)
	}

	// Watch for events in the directory
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				// If a write event occurs (file saved), send Rust code file to server
				sendRustCode(conn, filepath.Join(dir, "embedded.rs"))
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}

func sendRustCode(conn net.Conn, filePath string) {
	// Open the Rust code file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send the file over the connection
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}

	fmt.Println("Rust code file sent successfully!")
}
