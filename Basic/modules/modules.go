package greetings

//Create a module -- Write a small module with functions you can call from another module.

import (
	"errors"
	"fmt"
	"time"
)

func Hello(name string) (string, error) {
	//Replace empty to find failing tests
	if name == "" {
		return "", errors.New("empty string error")
	}
	message := fmt.Sprintf("hi %v welcome", name)
	return message, nil
}

//go routines

func PrintAsync(msg string) string {
	go func() {
		time.Sleep(1 * time.Second)
		msg := msg + "HI"
		fmt.Println("Async go routine", msg)
	}()
	return msg

}
