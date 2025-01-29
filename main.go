package main

import (
	"fmt"

	"example.com/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello, World!")
}