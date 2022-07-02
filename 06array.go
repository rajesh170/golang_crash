package main

import "fmt"

func main() {
	var numbers [3]string

	numbers[0] = "one"
	numbers[1] = "two"
	numbers[2] = "three"

	fmt.Println(numbers)

	var colors = [3]string{"red", "green", "blue"}
	fmt.Println(colors)
	fmt.Println(colors[2])
	fmt.Println(len(colors))

}
