package main

import "fmt"

func main() {
	score := make(map[string]int)
	score["aashish"] = 89
	fmt.Println(score)

	score["ram"] = 40
	score["shyam"] = 45
	score["hari"] = 49
	fmt.Println(score)

	getRamScore := score["ram"]
	fmt.Println(getRamScore)

	delete(score, "shyam")
	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v \n", k, v)
	}
}
