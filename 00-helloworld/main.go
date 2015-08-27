// File used to practice the basics of Go.

package main

// Including libraries
import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
	// Remote libraries
)

const Pi = 3.14

// Declaring and initialiting several vars:
var c, python, golang bool = false, false, true

// Structs
type Vertex struct {
	X float64
	Y float64
}

func getOS() string {
	currentOS := ""
	switch os := runtime.GOOS; os {
	case "darwin":
		currentOS = "OS X"
	case "linux":
		currentOS = "Linux"
	default:
		// freebsd, openbsd,windows...
		currentOS = os
	}

	return currentOS
}
func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Running under:", getOS())
}
