package models

import "github.com/nu7hatch/gouuid"

type Match struct{
    Id string
    History *History
}

func NewMatch() *Match {
    id, _ := uuid.NewV4()
    idString := id.String()

    return &Match{
        Id : idString,
        History: &History{},
    }
}