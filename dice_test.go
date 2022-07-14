package main

import (
	"reflect"
	"testing"
)

func TestRoll_Add(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	roll.Add(Roll{Kill: 1, Shield: 4})
	expect(t,
		reflect.DeepEqual(roll, Roll{Kill: 2, Push: 3, Shield: 4}))
}

func TestRoll_Add_Empty(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	roll.Add(Roll{})
	expect(t,
		reflect.DeepEqual(roll, Roll{Kill: 1, Push: 3}))
}

func TestRoll_Copy(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	copy := roll.Copy()
	delete(roll, Kill)
	expect(t,
		reflect.DeepEqual(roll, Roll{Push: 3}) &&
			reflect.DeepEqual(copy, Roll{Kill: 1, Push: 3}))
}

func TestDice_RollN_0(t *testing.T) {
	expect(t,
		reflect.DeepEqual(RedDice().RollN(0), Roll{}))
}

func BenchmarkDice_RollN(b *testing.B) {
	RedDice().RollN(6000)
}

func BenchmarkDiceSet_Roll(b *testing.B) {
	DiceSet{RedDice(): 6000}.Roll()
}
