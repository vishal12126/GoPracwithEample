// Go program to illustrate how to
// use for loop in the channel

package main

// Importing fmt, io and os
import (
	"fmt"
	"io"
	"os"
)

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string)

	// Anonymous goroutine
	go func() {
		mychnl <- "GFG"
		mychnl <- "gfg"
		mychnl <- "Test"
		mychnl <- "test123"
		close(mychnl)
	}()

	// Using for loop
	for res := range mychnl {
		fmt.Println(res)
	}

	const name, dept = "Test", "CS"

	// Calling Sprintf() function
	s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)
}
