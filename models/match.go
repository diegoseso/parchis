package models

import (
    "github.com/nu7hatch/gouuid"
    "fmt"
)

type Match struct{
    Id string
    MatchStatus *MatchStatus
    //History *History
}

type MatchStatus struct{
    BoardPicture map[int]byte
}

func NewMatchStatus() *MatchStatus{
    // Here we are gonna make a bytes map
    bs := make(map[int]byte)

    for i:=1;i<=4; i++{
        bs[1] = 0
    }

    return &MatchStatus{BoardPicture: bs }
}

func (ms *MatchStatus)CalculatePosition( player int, toAddValue int)int{

    if int(ms.BoardPicture[player]) + toAddValue > 75 {
        // This is the user bouncing
        ms.BoardPicture[player] = byte(75 - (int(ms.BoardPicture[player]) + toAddValue - 75))
        fmt.Println("Oh my god, it bounced")
        return int(ms.BoardPicture[player])
    }
    ms.BoardPicture[player] = ms.BoardPicture[player] + byte(toAddValue)
    return int(ms.BoardPicture[player])
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

