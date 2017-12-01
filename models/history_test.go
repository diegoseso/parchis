package models

import "time"

type History struct{
	StartDate time.Time
	EndDate   time.Time
	ActionLog []string
}