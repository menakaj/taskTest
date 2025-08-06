package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	log.Println(os.Environ())
	for i, arg := range os.Args[1:] {
		fmt.Println("Arg", i, ": ", arg)
	}
}
