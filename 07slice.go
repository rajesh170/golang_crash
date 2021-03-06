package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"red", "blue", "white"}
	fmt.Print(things)
	things = append(things, "green")
	fmt.Println(things)
	things = append(things[1 : len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "superman"
	heros[2] = "spiderman"
	fmt.Println(heros)
	heros = append(heros, "deadpool")
	heros = append(heros, "batman")
	fmt.Println(heros)
	fmt.Println(cap(heros))

	myInts := []int{4, 2, 6, 5, 8}

	isSorted := sort.IntsAreSorted(myInts)
	fmt.Println("Are ints sorted:", isSorted)
	sort.Ints(myInts)
	isSort := sort.IntsAreSorted(myInts)
	fmt.Println("Are ints sorted:", isSort)
	fmt.Println(myInts)
}
