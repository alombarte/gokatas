package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

const Pi = 3.14

// Declaring and initialiting several vars:
var c, python, golang bool = false, false, true

// Structs
type Vertex struct {
	X float64
	Y float64
}

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
}
