package models

type Player struct {
	Username string
	IsPlaying bool
}

type SessionID string

var OnlinePlayers map[string]*Player

func getSession(username string)(*SessionID, *Player, error){
	return nil, nil, nil
}

func setSession(username string)(*SessionID, *Player, error){
	return nil, nil, nil
}

func LoginUser(username string)(*SessionID, *Player, error){

    var sessionID *SessionID
	var player *Player
    sessionID, player, err := getSession(username)
    if err != nil{
    	sessionID, player, err = setSession(username)
    		if err != nil{
    			//we have to handle this error
			}

	}
	return sessionID, player, nil
}
