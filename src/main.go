package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	debugLogger *log.Logger
	errorLogger *log.Logger
	infoLogger  *log.Logger
)

// Function to initialize  different logging types (DEBUG/INFO) and formatting options
func initializeLogger() {
	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
}

// main
func main() {
	initializeLogger()

	// Load command-line flags configurations
	config := ParseConfigFromCommandLine()

	// Define a handler function to handle incoming requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // Write the response to the client
	}

	// Register the handler function to handle all requests to the root URL path "/"
	http.HandleFunc("/", handler)

	// Webserver takes the port option as a string, ex "":8080"
	portStr := ":" + strconv.Itoa(config.Port)

	// Start the web server
	infoLogger.Printf("Server is listening on port %d...\n", config.Port)
	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		errorLogger.Println(err)
	}
}
