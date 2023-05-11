package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define command-line flags
	name := flag.String("name", "World", "name to greet")
	flag.Parse()

	// Print the greeting
	message := fmt.Sprintf("Hello, %s!", *name)
	fmt.Println(message)
}
