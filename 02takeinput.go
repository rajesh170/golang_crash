package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var myString string
	// fmt.Scanln(&myString)
	// fmt.Println("your name is", myString)

	// var name string = "Rajesh"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your name:")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println("Your name is ", myname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating")
	myrating, _ := reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumrating + 2)
}
