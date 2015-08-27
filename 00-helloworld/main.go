// File used to practice the basics of Go.

package main

// Including libraries
import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

const Pi = 3.14159265358979323846264338327950288419716939937510582097494459

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

// swapText returns 2 strings
func swapText(x, y string) (string, string) {
	return y, x
}

// WordRepetition returns a map with the number of occurrences of each word inside a string.
func WordRepetition(s string) map[string]int {

	// Map declaration:
	var repetitions map[string]int
	repetitions = make(map[string]int) // After declaring a map needs to be initialized with 'make' before using it.

	words := strings.Fields(s)

	// Iterate using key => value in a range. Underscore means the var won't be used.
	for _, word := range words {
		repetitions[word]++
	}

	return repetitions
}

// fibonacci function returns an anonymous function that returns an integer itself.
func fibonacci() func() int {
	// fmt.Println("This code is executed only the first time")
	// Notice that this vars will be modified by the following anonymous function every time.
	prev, next := 0, 1

	return func() int {
		// fmt.Println("This code is executed subsequent times")
		prev, next = next, prev+next
		return prev
	}

}

func main() {

	// Mission accomplished:
	fmt.Println(swapText("world", "hello"))
	fmt.Println("---")

	// Printing: https://golang.org/pkg/fmt/
	fmt.Printf("π=%.5f (type %T).\n", Pi, Pi) // Print 5 decimals and variable type
	fmt.Println("The time is:", time.Now())
	fmt.Println("My favorite number is always:", rand.Intn(10)) // Note: deterministic operation
	fmt.Println("Running under:", getOS())
	fmt.Println("---")

	v := Vertex{3.0, -4.0}
	fmt.Println("Vertex:", v.X, v.Y, "Abs():", v.Abs())
	fmt.Println("---")

	a := "Am Zehnten Zehnten um zehn Uhr zehn zogen zehn zahme Ziegen zehn Zentner Zucker zum Zoo"
	fmt.Printf("Word repetition for string:\n%q\n", a)
	fmt.Println(WordRepetition(a))
	fmt.Println("---")

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, f())
	}
	fmt.Println("---")

}
