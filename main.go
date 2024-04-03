package main

import (
	"net"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Get the path from the user's argument (directory or file)
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// Check if the path points to a directory or a file
	info, err := os.Stat(path)
	if err != nil {
		http.Error(w, "File or directory not found", http.StatusNotFound)
		return
	}

	if info.IsDir() {
		// Serve files from the specified directory
		http.FileServer(http.Dir(path)).ServeHTTP(w, r)
	} else {
		// Serve the specific file
		http.ServeFile(w, r, path)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)

	// Find an available port and start the server
	for port := 8000; port < 9000; port++ {
		addr := fmt.Sprintf("127.0.0.1:%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			fmt.Printf("Server listening on port %d\n", port)
			cmd := exec.Command("open", fmt.Sprintf("http://localhost:%d", port))
			cmd.Run()

			fmt.Println("Press Ctrl+C to stop the server.")
			http.Serve(listener, nil)
			break
		}
	}
}

