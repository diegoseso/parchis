package models

import (
	"math/rand"
	"time"
)

const DICE_SIDES = 6

type Dice struct {
	value int
}

func NewDice() *Dice {
	return &Dice{}
}

func (d *Dice) Shake() int {

	now := time.Now()
	nanoSeed := now.UnixNano()

	rand.Seed(nanoSeed)
	d.value = rand.Intn(DICE_SIDES)
	return d.value
}

func (d *Dice) GetValue() int {
	return d.value
}
