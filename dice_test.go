package main

import (
	"reflect"
	"testing"
)

var roll = Roll{Kill: 1, Push: 3}

func TestRollAdd(t *testing.T) {
	roll.Add(Roll{Kill: 1, Shield: 4})
	if !reflect.DeepEqual(roll, Roll{Kill: 2, Push: 3, Shield: 4}) {
		t.Error()
	}
}

func TestRollAddEmpty(t *testing.T) {
	roll.Add(Roll{})
	if !reflect.DeepEqual(roll, roll) {
		t.Error()
	}
}

func TestDiceRollN0(t *testing.T) {
	if !reflect.DeepEqual(RedDice().RollN(0), Roll{}) {
		t.Error()
	}
}

func BenchmarkDiceRollN(b *testing.B) {
	RedDice().RollN(6000)
}

func BenchmarkDiceSetRoll(b *testing.B) {
	DiceSet{RedDice(): 6000}.Roll()
}
