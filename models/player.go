package models

import (
	"errors"
	"time"
	"github.com/davecgh/go-spew/spew"
)

const OnlinePlayersUpdate = 5 * time.Second

type Player struct {
	Username string
	IsPlaying bool
}

type SessionID string

var OnlinePlayers map[SessionID]*Player


func init(){
	OnlinePlayers = make(map[SessionID]*Player)
	spew.Dump(OnlinePlayers)
}

func getSession(username string)(SessionID, *Player, error){
	for session, player := range OnlinePlayers {
        if session == SessionID(username){
            return session, player, nil
	    }
	}
	return SessionID(""), &Player{}, errors.New("session does not exist")
}

func setSession(username string)(SessionID, *Player){
	OnlinePlayers[SessionID(username)] = &Player{Username: username, IsPlaying: true}
	return SessionID(username), OnlinePlayers[SessionID(username)]
}

func LoginUser(username string)(SessionID, *Player){
    sessionID, player, err := getSession(username)
    if err != nil{
    	sessionID, player = setSession(username)
	}
	return sessionID, player
}

func GetOnlinePlayers()[]string{
	onlinePlayers := make([]string, 0)
	for _, player := range OnlinePlayers {
        onlinePlayers = append(onlinePlayers, player.Username)
	}
	return onlinePlayers
}

type Players struct{
	Players []string `json:players`
}

type OnlinePLayersMsg struct{
	Type string         `json:"type"`
	Data *Players       `json:"data"`
}