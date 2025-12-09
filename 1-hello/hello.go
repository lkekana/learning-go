// https://go.dev/doc/tutorial/getting-started

package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

// used `go mod tidy` to install dependencies
// used `go run .` to run the program