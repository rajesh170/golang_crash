package main

import "fmt"

func main() {
	// var p *int
	// if p != nil {
	// 	fmt.Printf("p is having value: %v", *p)
	// } else {
	// 	fmt.Println("watch out for nil value")
	// }

	var lifebooster float64 = 99.2
	lifeboosterRef := &lifebooster
	ref := &lifeboosterRef
	fmt.Println(lifebooster)
	fmt.Println(lifeboosterRef)
	fmt.Println(*lifeboosterRef)
	fmt.Println(ref)

}
