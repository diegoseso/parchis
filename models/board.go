package models

import (
	"errors"
)

/**
 * Normal Board for 2 to 4 players with 4 Jails (One jail per player)
 */

/**
 *
 */
const BOXES = 76
const JAILS = 4
const TOKENS_PER_PLAYER = 4
const OPENING_VALUE = 5

type Board struct {
	Players []Player
	Match  *Match
	//Status *Status
}

func NewBoard() *Board{
	return &Board{}
}

func (B *Board)SetPlayer(Player Player)error{
	if len(B.Players) >= JAILS {
		return errors.New("No more players allowed")
	}
	B.Players = append(B.Players, Player)
	return nil
}

func (B *Board)StartNewMatch() {
		B.Match = NewMatch()
}
