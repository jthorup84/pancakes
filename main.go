package main

import (
	"fmt"
	"pancakes/pancakes"
)

func main() {
	input := "5\n-\n-+\n+-\n+++\n--+-"

	fmt.Println(input)

	output := pancakes.GetFlipCounts(input)

	fmt.Println(output)
}
