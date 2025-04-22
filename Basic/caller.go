package main

import (
	"fmt"

	greetings "Go-practice/Basic/modules"

	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	msg, err := greetings.Hello("Hello")

	ms := greetings.PrintAsync("HI")

	fmt.Println(ms)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
