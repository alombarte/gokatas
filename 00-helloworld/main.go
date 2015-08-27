// File used to practice the basics of Go.

package main

// Including libraries
import (
	"fmt"
	"math"
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

// Methods on structs. This is like Classes on other languages: V.abs()
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// getOS returns the current operattive system on which the program runs.
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
	fmt.Println("My favorite number is", rand.Intn(10)) // Caution: deterministic
	fmt.Println("Running under:", getOS())
	fmt.Println("---")

	v := Vertex{3.0, -4.0}
	fmt.Println("Vertex:", v.X, v.Y, "Abs():", v.Abs())

}
