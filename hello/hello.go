package main

import (
    "fmt"

	"github.com/lirajohny/GoStudy/tree/main/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
