package main

import (
	"os"
)

// Example_main ensures the following will work:
//
//	go build sumsquared.go
//	./sumsquared 2 3
func Example_main() {

	// Save the old os.Args and replace with our example input.
	oldArgs := os.Args
	os.Args = []string{"sumsquared", "2", "3"}
	defer func() { os.Args = oldArgs }()

	main()

	// Output:
	// (2 + 3) * (2 + 3) = 25
}
