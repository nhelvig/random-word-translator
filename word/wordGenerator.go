package main

import (
	"fmt"
	"random-word-translator/generator"
)

func main() {
	fmt.Printf(generator.GenerateRandomWord() + "\n")
}