package main

import (
	"fmt"

	"github.com/anurzhanuly/shiny-umbrella/array"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(
		array.ArrayStringsAreEqual(
			[]string{"abc", "d", "defg"},
			[]string{"a", "b", "c", "d", "def"}))
}
