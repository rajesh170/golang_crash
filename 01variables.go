package main

import "fmt"

func main() {
	var batman string = "I am batman"

	fmt.Println(batman)

	var superman string = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	// %T is to know the data type
	// thorRating := 3
	// fmt.Printf("%v,%T", thorRating, thorRating)

	var ironMan, catAmerica string = "I am ironman", "i am cat america"
	fmt.Println(ironMan, catAmerica)

	var defaultValue string
	fmt.Print(defaultValue)

	var (
		spiderman = "I am spiderman"
		age       = 18
		power     = "web slings,spidy senses"
		antman    = "i am antman"
	)
	fmt.Println(spiderman, age, power, antman)
}
