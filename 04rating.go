package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ = reader.ReadString('\n')

	//take user rating and convert into float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Rate dosa center between 1 to 5")
	userRating, _ = reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	fmt.Printf("Hello %v, \n Thanks for rating our Dosa center with %v star rating.\n \n Your rating was recorded in our system at %v \n \n", name, mynumrating, time.Now().Format(time.Stamp))

	if mynumrating == 5 {
		fmt.Println("Bonus for team for 5 star service")
	} else if mynumrating == 4 || mynumrating == 3 {
		fmt.Println("We are improving")
	} else if mynumrating < 3 {
		fmt.Println("We need to get serious now.")
	}
}
