package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	rob := User{"rob", "rob@gmal.com", 20}
	fmt.Printf("%+v \n", rob)
	fmt.Printf("%v \n", rob.Email)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@gmail.com"
	sam.Age = 23
	fmt.Printf("%+v \n", sam)

	tobby := &User{"tobby", "tobby@gmail.com", 27}
	fmt.Printf("%+v \n", tobby)
}
