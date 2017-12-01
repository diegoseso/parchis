package models

/**
 * Normal Board for 2 to 4 players with 4 Jails (One jail per player)
 */

/**
 *
 */
const BOXES = 76
const JAILS = 4
const TOKENS_PER_PLAYER = 4

type Board struct {
	Players []Player
	Match  *Match
	Status *Status
}

func NewBoard() *Board{
	return &Board{}
}

func (B *Board)SetPlayer(Player Player)error{
	if len(B.Players) >= JAILS {
		return nil
	}
	B.Players = append(B.Players, Player)
	return nil
}

func (B *Board)StartNewMatch() {
	if B.Status.IsReady() {
		B.Match = NewMatch()
		B.Status.Playing()
	}
}