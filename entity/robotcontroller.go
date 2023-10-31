package entity

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/util"
)

type RobotController struct{
	robot *Robot
	moving Mover
	facing Facer
}

func NewRobotController(xAxis, yAxis int, face string) *RobotController{	
	newRobot := NewRobot(event.NewCordinate(xAxis,yAxis), event.NewFacing(face))
	return &RobotController{robot: newRobot}
}

func (rc *RobotController) face(facing Facer){
	rc.facing = facing
	rc.facing.Face(rc.robot.facing)
}

func (rc *RobotController) move(moving Mover){
	rc.moving = moving
	rc.moving.Move(rc.robot)
}

func (rc *RobotController) doCommand(cmd string){
	if util.IsFacingCommand(cmd){
		if cmd == "R"{
			rc.face(&FaceRight{})
		}else if cmd == "L"{
			rc.face(&FaceLeft{})
		}
	}else if util.IsMoveCommand(cmd){
		if cmd == "A"{
			rc.move(&MoveAdvance{})
		}
	}
}

func (rc *RobotController) MovementsCommand(moves string)(xAxis, yAxis int, face string) {
	for _, move := range moves {
		rc.doCommand(string(move))
	}
	switch rc.robot.facing.Direction(){
	case 1:
		face = "N"
	case 2:
		face = "E"
	case 3:
		face = "S"
	case 4:
		face = "W"
	}
	return rc.robot.coordinate.XAxis(), rc.robot.coordinate.YAxis(), face
}


