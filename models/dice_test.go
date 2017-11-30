package models

import(
    "github.com/davecgh/go-spew/spew"
    "testing"
)


func TestDice(t *testing.T) {
	dice := NewDice()
	value := dice.Shake()
	spew.Dump(value)
}
