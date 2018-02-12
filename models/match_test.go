package models

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewMatchStatus(t *testing.T) {
	ms := NewMatchStatus()
	spew.Dump(ms)
}

func TestMatchStatus_CalculatePosition_Bouncing(t *testing.T) {

	match := NewMatch()

	bounce := match.MatchStatus.CalculatePosition(1, 77)
	if bounce != 73 {
		t.Fatal("The bouncing is not working ok")
	}
}

func TestMatchStatus_CalculatePosition_Winning(t *testing.T) {

	match := NewMatch()

	match.MatchStatus.CalculatePosition(1, 5)
	arrival := match.MatchStatus.CalculatePosition(1, 70)
	if arrival != 75 {
		t.Fatal("The arrival is not working ok, probably cause the bouncing")
	}
}

func TestMatchStatus_IsAWinner(t *testing.T) {
	match := NewMatch()

	match.MatchStatus.BoardPicture[1][1] = byte(BOXES)
	match.MatchStatus.BoardPicture[1][2] = byte(BOXES)
	match.MatchStatus.BoardPicture[1][3] = byte(BOXES)
	match.MatchStatus.BoardPicture[1][4] = byte(BOXES)

	if !match.MatchStatus.IsAWinner(1) {
		t.Fatal("Is not calculating the winner well...")
	}

}

func TestMatchStatus_GetMovablePawns(t *testing.T) {
	match := NewMatch()

	match.MatchStatus.BoardPicture[1] = PlayerStatus{1: 0, 2: 1, 3: 1, 4: 0}

	if len(match.MatchStatus.GetMovablePawns(1)) != 2 {
		t.Fatal("We are not getting the movable pawns right")
	}
}

func TestMatchStatus_GetPawnFromJail(t *testing.T) {
	match := NewMatch()

	_, err := match.MatchStatus.GetPawnFromJail(1)
	if err != nil {
		t.Fatal("a Riot has happened at Jail")
	}
}

func TestCompleteMatchStandAlone(t *testing.T) {
	player1 := Player{Username: "one"}
	player2 := Player{Username: "two"}
	player3 := Player{Username: "three"}
	player4 := Player{Username: "four"}

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
	testLoopControl := 0

	for testLoopControl <= 2000 && endGame == false {
		testLoopControl++
		t.Log("Got into the playing loop")

		// @TODO Loop through all players, right now it considers there are always 4 players...
		if control == JAILS+1 {
			control = 1
		}

		dice.Shake()
		t.Logf("%v", dice.GetValue())
		t.Logf("Player %v got a %v", control, dice.GetValue())

		if dice.GetValue() == OPENING_VALUE && board.Match.MatchStatus.IsThereAPawnInJail(int(control)) {

			// Get a Pawn out of the Jail into Play
			pawnIndex, err := board.Match.MatchStatus.GetPawnFromJail(int(control))
			if err != nil {
				t.Fatal("There is a hell of an error on the pawns")
			}
			board.Match.MatchStatus.BoardPicture[int(control)][pawnIndex] = 1

		} else {

			// Make a move with the dice result
			//for pawn := range board.Match.MatchStatus.GetMovablePawns(int(control)) {
			board.Match.MatchStatus.CalculatePosition(int(control), dice.GetValue())
			//	break
			//}
		}

		if board.Match.MatchStatus.IsAWinner(int(control)) {
			t.Logf("Player%v wins", control)
			t.Log(board.Match.MatchStatus.BoardPicture)
			endGame = true
		}

		t.Log(board.Match.MatchStatus.BoardPicture)
		control++
	}

}
