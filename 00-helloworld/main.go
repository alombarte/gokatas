// File used to practice the basics of Go.

package main

// Including libraries
import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

const Pi = 3.14159265358979323846264338327950288419716939937510582097494459

// Declaring and initialiting several vars:
var c, python, golang bool = false, false, true

// Structs
type Band struct {
	Name  string
	Genre string
	Year  int
}

// Age returns how old the band is in years.
// Example of methods on structs. This is like Classes on other languages: b.Age()
func (b *Band) Age() int {
	return time.Now().Year() - b.Year
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

	// Mission accomplished!:
	fmt.Println(swapText("world", "hello"))
	fmt.Println("---")

	// Printing: https://golang.org/pkg/fmt/
	fmt.Printf("π=%.5f (type %T).\n", Pi, Pi) // Print 5 decimals and variable type
	fmt.Println("The time is:", time.Now())
	fmt.Println("My favorite number is always:", rand.Intn(10)) // Note: deterministic operation
	fmt.Println("Running under:", getOS())
	fmt.Println("---")

	fmt.Println("Basics...")

	// Pointers:
	i, j := 10, 50
	p := &i         // p points to i
	fmt.Println(*p) // read i through the pointer (prints 10)
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i (prints 21)

	p = &j         // point to j
	*p = *p / 2    // divide j through the pointer
	fmt.Println(j) // see the new value of j (25)
	fmt.Println("---")

	// Declare an array of integers
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	//Loop
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	//Convenient loop using range (key => value)
	for k, v := range s {
		fmt.Printf("s[%d] = %d, ", k, v)
	}

	fmt.Printf("\nValues: ")
	//Loop using range and ignoring key or value (place an underscore)
	for _, v := range s {
		fmt.Printf("%d, ", v)
	}
	fmt.Println("---")

	// Slices
	slice := make([]int, 5)
	// Grow it over its capacity with append
	slice = append(slice, 5, 6, 7, 8, 9, 10)
	fmt.Println("Slice:", slice, "Len:", len(slice), "Capacity:", cap(slice))
	fmt.Println("---")

	// Using structs
	b := Band{"Metallica", "Heavy metal", 1981}
	fmt.Println(b, "Age:", b.Age())

	// Map of struct Band
	var bands map[string]Band
	bands = make(map[string]Band)

	bands["RHCP"] = Band{"Red Hot Chili Peppers", "Heavy metal", 1983}
	bands["Metallica"] = Band{"Metallica", "Heavy metal", 1981}
	fmt.Println(bands["Metallica"].Year, bands["Metallica"].Genre)

	// Or use map literals:
	var m2 = map[string]Band{
		"RHCP": Band{
			"Red Hot Chili Peppers", "Heavy metal", 1983,
		}, // Type "Band" can be omitted if only using one:
		"Metallica": {
			"Metallica", "Heavy metal", 1981,
		},
	}

	fmt.Println(m2["RHCP"].Name, m2["joselito"].Genre)

	//Map mutation
	delete(m2, "Metallica")
	_, key_exists := m2["Metallica"]
	if !key_exists {
		fmt.Println("Map mutated correctly")
	}

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

	// Defering executions. When you defer sonething gets stacked until the end of the execution:
	defer fmt.Println(">>...WORLD! (1st deferred)")
	defer fmt.Println(">>...2nd deferred")
	defer fmt.Println(">>...3rd deferred")

	fmt.Println(">>GOODBYE...")

}
