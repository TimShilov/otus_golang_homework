package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	message := "Hello, OTUS!"
	reversedMessage := stringutil.Reverse(message)
	fmt.Println(reversedMessage)
}
