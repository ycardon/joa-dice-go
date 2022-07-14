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

func TestRoll_Cancel_NoFace(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	shieldCount := 3
	roll.cancel(Disrupt, &shieldCount)
	expect(t,
		shieldCount == 3 &&
			reflect.DeepEqual(roll, Roll{Kill: 1, Push: 3}))
}
func TestRoll_Cancel_0(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	shieldCount := 0
	roll.cancel(Push, &shieldCount)
	expect(t,
		shieldCount == 0 &&
			reflect.DeepEqual(roll, Roll{Kill: 1, Push: 3}))
}

func TestRoll_Cancel_Smaller(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	shieldCount := 2
	roll.cancel(Push, &shieldCount)
	expect(t,
		shieldCount == 0 &&
			reflect.DeepEqual(roll, Roll{Kill: 1, Push: 1}))
}
func TestRoll_Cancel_Equal(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	shieldCount := 3
	roll.cancel(Push, &shieldCount)
	expect(t,
		shieldCount == 0 &&
			reflect.DeepEqual(roll, Roll{Kill: 1}))
}

func TestRoll_Cancel_Bigger(t *testing.T) {
	roll := Roll{Kill: 1, Push: 3}
	shieldCount := 4
	roll.cancel(Push, &shieldCount)
	expect(t,
		shieldCount == 1 &&
			reflect.DeepEqual(roll, Roll{Kill: 1}))
}

func TestRoll_ApplyDefense(t *testing.T) {
	att := Roll{Kill: 1, Disrupt: 2, Push: 3, Shield: 4, Blank: 5}
	def := Roll{Shield: 5, Push: 10}
	res := ResolveAttack(att, def)
	expect(t,
		reflect.DeepEqual(res, Roll{Push: 1}))
}

func TestRoll_ApplyDefense_CompleteCancel(t *testing.T) {
	att := Roll{Kill: 1, Disrupt: 2, Push: 3, Shield: 4, Blank: 5}
	def := Roll{Shield: 6, Push: 10}
	res := ResolveAttack(att, def)
	expect(t,
		reflect.DeepEqual(res, Roll{}))
}

func TestRoll_ApplyDefense_None(t *testing.T) {
	att := Roll{Kill: 1, Disrupt: 2, Push: 3, Shield: 4, Blank: 5}
	def := Roll{Push: 10}
	res := ResolveAttack(att, def)
	expect(t,
		reflect.DeepEqual(res, Roll{Kill: 1, Disrupt: 2, Push: 3}))
}
