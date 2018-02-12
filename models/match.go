package models

import (
	"errors"
	"github.com/nu7hatch/gouuid"
)

type Match struct {
	Id          string
	MatchStatus *MatchStatus
	//History *History
}

func NewMatch() *Match {
    id, _ := uuid.NewV4()
    idString := id.String()

    return &Match{
        Id:          idString,
        MatchStatus: NewMatchStatus(),
    }
}

type TurnControl int

type MatchStatus struct {
	BoardPicture map[int]PlayerStatus
}

func NewMatchStatus() *MatchStatus {
    // Here we are gonna make a bytes map
    bs := make(map[int]PlayerStatus)

    for i := 1; i <= 4; i++ {
        bs[i] = PlayerStatus{1: 0, 2: 0, 3: 0, 4: 0}
    }

    return &MatchStatus{BoardPicture: bs}
}

func (ms *MatchStatus) CalculatePosition(player int, toAddValue int) int {

    // @TODO quite a challenge get the algorithm for the best choice

    for k, v := range ms.BoardPicture[player] {
        if int(v)+toAddValue > 75 {

            ms.BoardPicture[player][k] = byte(75 - (int(v) + toAddValue - 75))

            return int(ms.BoardPicture[player][k])
        }
        ms.BoardPicture[player][k] += byte(toAddValue)
        return int(ms.BoardPicture[player][k])
    }
    // @TODO This is gonna be a bug one day
    return 0
}

func (ms *MatchStatus) IsAWinner(player int) bool {
    for _, pawnPosition := range ms.BoardPicture[player] {
        if pawnPosition != BOXES {
            return false
        }
    }
    return true
}

func (ms *MatchStatus) IsThereAPawnInJail(player int) bool {
    for _, pawnPosition := range ms.BoardPicture[player] {
        if pawnPosition == 0 {
            return true
        }
    }
    return false
}

func (ms *MatchStatus) GetPawnFromJail(player int) (int, error) {
    for k, pawnPosition := range ms.BoardPicture[player] {
        if pawnPosition == 0 {
            return k, nil
        }
    }
    return 0, errors.New("Not Pawn left in Jail")
}

func (ms *MatchStatus) GetMovablePawns(player int) []int {
    var freePawns []int
    for k, pawnPosition := range ms.BoardPicture[player] {
        if pawnPosition != 0 && pawnPosition != BOXES {
            freePawns = append(freePawns, k)
        }
    }
    return freePawns
}

type PlayerStatus map[int]byte



