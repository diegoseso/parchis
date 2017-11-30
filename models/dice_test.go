package models

func TestDice(t *testing.T) {
   dice := models.NewDice()
   value := dice.Shake()
   spew.Dump(value)
}
