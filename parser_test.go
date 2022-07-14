package main

import (
	"reflect"
	"testing"
)

func TestParse_Attack(t *testing.T) {
	att, _, _ := Parse("1R")
	expect(t,
		reflect.DeepEqual(att, DiceSet{RedDice(): 1}))
}

func TestParse_AttackMultiple(t *testing.T) {
	att, _, _ := Parse("1R 2R")
	expect(t,
		reflect.DeepEqual(att, DiceSet{RedDice(): 3}))
}

func TestParse_Empty(t *testing.T) {
	att, _, _ := Parse("")
	expect(t,
		reflect.DeepEqual(att, DiceSet{}))
}

func TestParse_AttackDefence(t *testing.T) {
	att, def, isDef := Parse("1R - 2W")
	expect(t,
		reflect.DeepEqual(att, DiceSet{RedDice(): 1}) &&
			reflect.DeepEqual(def, DiceSet{WhiteDice(): 2}) &&
			isDef == true)
}

func TestParse_DefenceOnly(t *testing.T) {
	att, def, isDef := Parse("/ 2W")
	expect(t,
		reflect.DeepEqual(att, DiceSet{}) &&
			reflect.DeepEqual(def, DiceSet{WhiteDice(): 2}) &&
			isDef == true)
}
