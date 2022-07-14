package main

import (
	"math/rand"
	"time"
)

/********** FACE **********/

// a dice Face
type Face int

// pretty print a Face
func (f Face) String() string {
	return [...]string{"Kill", "Disrupt", "Push", "Shield", "Blank", "Trample", "Death", "Rally", "DelayedRally"}[f]
}

// (one of) the known dice Faces
const (
	Kill Face = iota
	Disrupt
	Push
	Shield
	Blank
	Trample
	Death
	Rally
	DelayedRally
)

/********** ROLL **********/

// the result of a dice Roll
type Roll map[Face]int

// add a Roll result to an existing one
func (r1 Roll) Add(r2 Roll) {
	for face := range r2 {
		r1[face] += r2[face]
	}
}

// cancel Roll face count by an amount of shield, return remaining amount
func (r Roll) Cancel(face Face, shieldAmount int) int {
	faceCount, ok := r[face]
	if !ok {
		return shieldAmount
	} else {
		if faceCount > shieldAmount {
			r[face] -= shieldAmount
			return 0
		} else {
			delete(r, face)
			return shieldAmount - faceCount
		}
	}
}

/********** DICE **********/

// a Dice with 6 faces
type Dice [6]Face

// roll 1 Dice
func (d Dice) Roll() Face {
	roll := rand.Intn(len(d))
	return Face(d[roll])
}

// roll N Dice
func (d Dice) RollN(n int) Roll {
	roll := Roll{}
	for i := 1; i <= n; i++ {
		face := d.Roll()
		roll.Add(Roll{face: 1})
	}
	return roll
}

// the known Dice
func BlackDice() Dice    { return Dice{Kill, Disrupt, Disrupt, Shield, Shield, Shield} }
func RedDice() Dice      { return Dice{Kill, Kill, Disrupt, Disrupt, Push, Shield} }
func YellowDice() Dice   { return Dice{Disrupt, Push, Push, Shield, Blank, Blank} }
func WhiteDice() Dice    { return Dice{Disrupt, Disrupt, Push, Shield, Shield, Blank} }
func GiganticDice() Dice { return Dice{Kill, Disrupt, Disrupt, Push, Trample, Trample} }
func DoomDice() Dice     { return Dice{Disrupt, Death, Death, Rally, Rally, DelayedRally} }

/********** DICE SET **********/

// a set of Dice
type DiceSet map[Dice]int

// roll a DiceSet
func (ds DiceSet) Roll() Roll {
	roll := Roll{}
	for dice, count := range ds {
		roll.Add(dice.RollN(count))
	}
	return roll
}

/********** INIT **********/

// set the random seed
func InitRandom() {
	rand.Seed(time.Now().UnixNano())
}
