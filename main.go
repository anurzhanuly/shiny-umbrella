package main

import "fmt"

func main() {
	index := 10
	result := 0

	for i := index; i > 0; i-- {
		fmt.Println("Superman")
		result += i
	}

	fmt.Println(result)
}
