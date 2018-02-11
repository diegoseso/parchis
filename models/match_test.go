package models

import (
	"testing"
	"github.com/davecgh/go-spew/spew"
)


func TestNewMatchStatus(t *testing.T) {
	ms := NewMatchStatus()
	spew.Dump(ms)
}

func TestMatchStatus_CalculatePosition(t *testing.T) {

	match := NewMatch()

	match.MatchStatus.CalculatePosition(1, 5)
	t.Log(match.MatchStatus.CalculatePosition(1, 72))
}

func TestCompleteMatchStandAlone(t *testing.T) {
	t.Log("Got into the playing loop")
	player1 := Player{Username:"one"}
	player2 := Player{Username:"two"}
	player3 := Player{Username:"three"}
	player4 := Player{Username:"four"}

    board := NewBoard()
    board.SetPlayer(player1)
	board.SetPlayer(player2)
	board.SetPlayer(player3)
	board.SetPlayer(player4)

	board.StartNewMatch()

	endGame := false

	t.Log("Getting dice prepared")
	dice := NewDice()
	control := TurnControl(1)

	for endGame != true {
		t.Log("Got into the playing loop")
		if control == JAILS + 1 {
			control = 1
		}

		s := dice.Shake()
		t.Logf("%v", s)
		t.Logf("%v", dice.GetValue())
		t.Logf("Player %v got a %v", control, dice.GetValue())

		if dice.GetValue() == OPENING_VALUE && board.Match.MatchStatus.BoardPicture[int(control)] == 0{
			board.Match.MatchStatus.BoardPicture[int(control)] = 1
		}
		if board.Match.MatchStatus.BoardPicture[int(control)] > 0{
			board.Match.MatchStatus.CalculatePosition(int(control), dice.GetValue())
		}

		if board.Match.MatchStatus.BoardPicture[int(control)] > 75 {
			t.Logf("Player%v wins", control)
			t.Log(board.Match.MatchStatus.BoardPicture)
			endGame = true
		}

		control++
	}


}
