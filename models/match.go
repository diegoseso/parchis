package models

import (
    "github.com/nu7hatch/gouuid"
    "fmt"
    "errors"
)

type Match struct{
    Id string
    MatchStatus *MatchStatus
    //History *History
}

type MatchStatus struct{
    BoardPicture map[int]PlayerStatus
}

type PlayerStatus map[int]byte

func NewMatchStatus() *MatchStatus{
    // Here we are gonna make a bytes map
    bs := make(map[int]PlayerStatus)

    for i:=1;i<=4; i++{
        bs[i] = PlayerStatus{1:0,2:0,3:0,4:0}
    }

    return &MatchStatus{BoardPicture: bs }
}

func (ms *MatchStatus)CalculatePosition( player int, toAddValue int)int{

    if int(ms.BoardPicture[player][1]) + toAddValue > 75 {
        // This is the user bouncing
        ms.BoardPicture[player][1] = byte(75 - (int(ms.BoardPicture[player][1]) + toAddValue - 75))
        fmt.Println("Oh my god, it bounced")
        return int(ms.BoardPicture[player][1])
    }
    ms.BoardPicture[player][1] = ms.BoardPicture[player][1] + byte(toAddValue)
    return int(ms.BoardPicture[player][1])
}

func (ms *MatchStatus)IsAWinner(player int)bool{
    for _, pawnPosition := range ms.BoardPicture[player]{
        if pawnPosition != BOXES{
            return false
        }
    }
    return true
}

func (ms *MatchStatus)IsThereAPawnInJail(player int)bool{
    for _, pawnPosition := range ms.BoardPicture[player]{
        if pawnPosition == 0{
            return true
        }
    }
    return false
}

func (ms *MatchStatus)GetPawnFromJail(player int)(int, error){
    for k, pawnPosition := range ms.BoardPicture[player]{
        if pawnPosition == 0{
            return k, nil
        }
    }
    return 0, errors.New("Not Pawn left in Jail")
}

func (ms *MatchStatus)GetMovablePawns(player int)[]int{
    var freePawns []int
    for k, pawnPosition := range ms.BoardPicture[player]{
        if pawnPosition != 0 && pawnPosition != BOXES{
            freePawns = append(freePawns, k)
        }
    }
    return freePawns
}


func NewMatch() *Match {
    id, _ := uuid.NewV4()
    idString := id.String()

    return &Match{
        Id : idString,
        MatchStatus: NewMatchStatus(),
    }
}

type TurnControl int

