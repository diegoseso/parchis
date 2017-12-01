package models

const AVAILABLE = "AVAILABLE"
const READY = "READY"
const PLAYING = "PLAYING"

type Status struct{
	Value string
}

func(S *Status)IsAvailable() bool{
	if S.Value == AVAILABLE {
		return true
	}
	return false
}

func(S *Status)IsReady() bool{
	if S.Value == READY {
		return true
	}
	return false
}

func(S *Status)IsPlaying() bool{
	if S.Value == PLAYING {
		return true
	}
	return false
}

func(S *Status)Playing() error{
	if S.Value == READY {
		S.Value = PLAYING
	}
}

func(S *Status)Available() error{
	S.Value = AVAILABLE
}

func(S *Status)Ready() error{
    if S.Value == PLAYING || S.Value == AVAILABLE {
		S.Value = READY
	}
}