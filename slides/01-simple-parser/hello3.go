package main

import "fmt"

type Greeter struct {
	Name  string
	close chan interface{}
}

func (g Greeter) greet(name string) {
	fmt.Printf("%s: Hello, %s!\n", g.Name, name)
	g.close <- nil
}

func main() {
	name := "World" // good enough for this example
	var greeter Greeter
	done := make(chan interface{})
	for i := 0; i < 10; i++ {
		greeter = Greeter{fmt.Sprintf("Greeter #%d", i), done}
		go greeter.greet(name)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	fmt.Println("We're done!")
}
