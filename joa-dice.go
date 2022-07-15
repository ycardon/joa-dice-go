package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/ryanuber/columnize"
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

// pretty print a Roll
func (r Roll) String() string {
	var s string
	for i := 0; i < int(FACE_MAX); i++ {
		value, ok := r[Face(i)]
		if ok && value > 0 {
			s += fmt.Sprintf("%d %s|", value, Face(i))
		}
	}
	if len(s) > 1 {
		s = s[:len(s)-1]
	}
	return s
}

// get a dice set from the CLI, roll and resolve it
func main() {
	InitRandom()

	// parse input from CLI
	input := strings.Join(os.Args[1:], " ")
	attackDiceSet, defenceDiceSet, isDef := Parse(input)

	// init colomnize
	config := columnize.DefaultConfig()
	config.Glue = color.InBlue(" | ")

	// roll and print the results
	attack := attackDiceSet.Roll()
	if !isDef {
		output := []string{fmt.Sprintf("%s|%s", color.InBold("attack"), attack)}
		fmt.Println(columnize.Format(output, config))
	} else {
		defence := defenceDiceSet.Roll()
		result := ResolveAttack(attack, defence)
		output := []string{
			fmt.Sprintf("%s|%s", color.InRed("attack"), attack),
			fmt.Sprintf("%s|%s", color.InGreen("defense"), defence),
			fmt.Sprintf("%s|%s", color.InCyan("result"), result),
		}
		fmt.Println(columnize.Format(output, config))
	}
}
