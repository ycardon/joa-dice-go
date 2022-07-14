package main

import (
	"strconv"
	"strings"
)

func Parse(s string) (attack DiceSet, defence DiceSet, isDef bool) {
	attack = make(DiceSet)
	defence = make(DiceSet)
	isDef = false

	words := strings.Fields(s)
	for _, word := range words {
		if word == "-" || word == "/" {
			isDef = true
		} else {
			init := word[:len(word)-1]
			last := strings.ToUpper(word)[len(word)-1]
			dice, ok := pareDice(last)
			if ok {
				if !isDef {
					attack[dice] += parseInt(init)
				} else {
					defence[dice] += parseInt(init)
				}
			}
		}
	}
	return
}

func parseInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		value = 1
	}
	return value
}

func pareDice(char byte) (d Dice, ok bool) {
	ok = true
	switch char {
	case 'B':
		d = BlackDice()
	case 'R':
		d = RedDice()
	case 'Y':
		d = YellowDice()
	case 'W':
		d = WhiteDice()
	case 'G':
		d = GiganticDice()
	case 'D':
		d = DoomDice()
	default:
		ok = false
	}
	return
}
