package main

import (
	"fmt"
	"os"
	"strings"
)

// resolve an attack
func ResolveAttack(attack Roll, defense Roll) Roll {
	result := attack.Copy()

	// apply defense shields on the attack
	shieldCount := defense[Shield]
	result.cancel(Kill, &shieldCount)
	result.cancel(Disrupt, &shieldCount)
	result.cancel(Push, &shieldCount)

	// remove unrelevant faces from the attack
	delete(result, Shield)
	delete(result, Blank)
	return result
}

// cancel Roll face by an amount of shield, return remaining amount
func (r Roll) cancel(face Face, shieldCount *int) {
	faceCount, ok := r[face]
	if ok {
		if faceCount > *shieldCount {
			r[face] -= *shieldCount
			*shieldCount = 0
		} else {
			delete(r, face)
			*shieldCount -= faceCount
		}
	}
}

// get a dice set from the CLI, roll and resolve it
func main() {

	InitRandom()
	format := "%s = %s\n"

	// parse input from CLI
	input := strings.Join(os.Args[1:], " ")
	attackDiceSet, defenceDiceSet, isDef := Parse(input)

	// roll and print the results
	attack := attackDiceSet.Roll()
	if !isDef {
		fmt.Printf(format, "attack", attack)
	} else {
		defence := defenceDiceSet.Roll()
		result := ResolveAttack(attack, defence)
		fmt.Printf(format, "attack", attack)
		fmt.Printf(format, "defense", defence)
		fmt.Printf(format, "result", result)
	}
}
