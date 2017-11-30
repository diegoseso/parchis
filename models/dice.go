package models

import("math/rand")

const DICE_SIDES=6

type Dice Struct{
    value  int
}

func NewDice()Dice{
   return Dice{}
}

func( d *Dice ) Shake (int){
    rand.Seed(42)
    d.value = rand.Intn(DICE_SIDES)
    return *d.value
}

func( d *Dice ) GetValue(int){
    return *d.Value
}
