package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"

type Robot struct {
	coordinate *event.Cordinates
	facing      *event.Facing
}

func NewRobot(cordinate *event.Cordinates, facing *event.Facing) *Robot {
	return &Robot{coordinate: cordinate, facing: facing}
}