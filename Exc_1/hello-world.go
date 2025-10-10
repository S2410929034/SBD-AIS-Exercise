package main

import "fmt"

type HelloWorld struct {
	Message string
	Number  int
}

func (h HelloWorld) Print() {
	fmt.Printf("%s | Number %d", h.Message, h.Number)
}

func main() {
	hw := HelloWorld{Message: "Hello, World!", Number: 1}
	hw.Print()
}
