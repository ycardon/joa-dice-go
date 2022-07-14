package main

import "fmt"

func main() {
	InitRandom()

	fmt.Println("roll 6000 times", RedDice())
	fmt.Println(RedDice().RollN(6000))

	attack, defence, _ := Parse("1R")
	fmt.Println("attack:", attack, "defense:", defence)
}
