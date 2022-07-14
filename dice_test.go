package main

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, ok bool) {
	if !ok {
		t.Error()
	}
}

func TestRollAdd(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	roll.Add(Roll{Kill: 1, Shield: 4})
	expect(t,
		reflect.DeepEqual(roll, Roll{Kill: 2, Push: 3, Shield: 4}))
}

func TestRollAddEmpty(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	roll.Add(Roll{})
	expect(t,
		reflect.DeepEqual(roll, roll))
}

func TestRollCancelNone(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	remain := roll.Cancel(Disrupt, 3)
	expect(t,
		remain == 3 && reflect.DeepEqual(roll, Roll{Kill: 1, Push: 3}))
}

func TestRollCancelSmaller(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	remain := roll.Cancel(Push, 2)
	expect(t,
		remain == 0 && reflect.DeepEqual(roll, Roll{Kill: 1, Push: 1}))
}
func TestRollCancelEqual(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	remain := roll.Cancel(Push, 3)
	expect(t,
		remain == 0 && reflect.DeepEqual(roll, Roll{Kill: 1}))
}

func TestRollCancelBigger(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	remain := roll.Cancel(Push, 4)
	expect(t,
		remain == 1 && reflect.DeepEqual(roll, Roll{Kill: 1}))
}

func TestDiceRollN0(t *testing.T) {
	expect(t,
		reflect.DeepEqual(RedDice().RollN(0), Roll{}))
}

func BenchmarkDiceRollN(b *testing.B) {
	RedDice().RollN(6000)
}

func BenchmarkDiceSetRoll(b *testing.B) {
	DiceSet{RedDice(): 6000}.Roll()
}
