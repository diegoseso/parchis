package models

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestDice(t *testing.T) {
	dice := NewDice()
	for i := 1; i < 100; i++ {
		value := dice.Shake()
		spew.Dump(dice.GetValue())
		spew.Dump(value)
	}
}
