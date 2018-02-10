package models

import "github.com/nu7hatch/gouuid"

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

func NewMatch() *Match {
    id, _ := uuid.NewV4()
    idString := id.String()

    return &Match{
        Id : idString,
        MatchStatus: NewMatchStatus(),
    }
}

type TurnControl int

