package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hi !", time.Now())
}

// To run the program, put the code in hello-world.go and use go run.

// $ go run basic/hi.go
// Hi !
// Sometimes we’ll want to build our programs into binaries. We can do this using go build.

// $ go build basic/hi.go
// $ ls
// basic/ hi
// We can then execute the built binary directly.

// $ ./hi
// Hi !
