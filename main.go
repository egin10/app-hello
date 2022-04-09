package main

import (
	"fmt"
	go_hello "github.com/egin10/go-hello"
)

func main() {
	hello := go_hello.Hello()
	fmt.Println(hello)

	sayHelloTo := go_hello.SayHelloTo("Egin")
	fmt.Println(sayHelloTo)
}
