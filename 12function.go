package main

import "fmt"

func main() {
	superman()
}

func superman() {
	fmt.Println("I am superman")
	result := multiplyme(6, 3)
	fmt.Println(result)
	myresult, lenth, name := adder(1, 4, 5, 6, 7)
	fmt.Println(myresult, lenth, name)
	mulresult, loc := multiplyer(4, 5, 3)
	fmt.Println(mulresult, loc)

}

func multiplyme(num1 int, num2 int) int {
	return num1 * num2
}

func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}

func multiplyer(values ...int) (int, string) {
	mul := 1
	for i := range values {
		mul = mul * values[i]
	}
	loc := "Multiplyer string"

	return mul, loc
}
