package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"

type Robot struct {
	coordinate *event.Cordinates
	facing      *event.Facing
}

func NewRobot(coordinate *event.Cordinates, facing *event.Facing) *Robot {
	return &Robot{coordinate: coordinate, facing: facing}
}

// func (r *Robot) addCordinate(xAxis, yAxis int) {
// 	r.xAxis = xAxis
// 	r.yAxis = yAxis
// }

func (r *Robot) Face() *event.Facing{
	return r.facing
}

// func (r *Robot) Move(movement event.Movement) {
// 	r.moving.move(movement)
// }
