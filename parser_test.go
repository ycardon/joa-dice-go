package main

import (
	"reflect"
	"testing"
)

func TestParseAttack(t *testing.T) {
	att, _, _ := Parse("1R")
	if !reflect.DeepEqual(att, DiceSet{RedDice(): 1}) {
		t.Error()
	}
}

func TestParseAttackMultiple(t *testing.T) {
	att, _, _ := Parse("1R 2R")
	if !reflect.DeepEqual(att, DiceSet{RedDice(): 3}) {
		t.Error()
	}
}

func TestParseEmpty(t *testing.T) {
	att, _, _ := Parse("")
	if !reflect.DeepEqual(att, DiceSet{}) {
		t.Error()
	}
}

func TestParseAttackDefence(t *testing.T) {
	att, def, isDef := Parse("1R - 2W")
	if !reflect.DeepEqual(att, DiceSet{RedDice(): 1}) &&
		!reflect.DeepEqual(def, DiceSet{WhiteDice(): 2}) &&
		isDef != true {
		t.Error()
	}
}

func TestParseDefenceOnly(t *testing.T) {
	att, def, isDef := Parse("/ 2W")
	if !reflect.DeepEqual(att, DiceSet{}) &&
		!reflect.DeepEqual(def, DiceSet{WhiteDice(): 2}) &&
		isDef != true {
		t.Error()
	}
}
