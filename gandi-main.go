package main

import (
	"fmt"

	// gandi_say_hello "github.com/PutuGandi/gandi-go-say-hello"
	gandi_say_hello "github.com/PutuGandi/gandi-go-say-hello/v2"
)

// func main() {
// 	fmt.Println(gandi_say_hello.SayHello())
// }

func main() {
	fmt.Println(gandi_say_hello.SayHello("Gandi Ganteng"))
}
