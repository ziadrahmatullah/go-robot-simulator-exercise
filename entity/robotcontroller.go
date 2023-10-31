package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"

type RobotController struct{
	robot *Robot
	moving Mover
	facing Facer
}

func NewRobotController(xAxis, yAxis int, face string) *RobotController{	
	newRobot := NewRobot(event.NewCordinate(xAxis,yAxis), event.NewFacing(face))
	return &RobotController{robot: newRobot}
}

func (rc *RobotController) Face(facing Facer){
	rc.facing = facing
	rc.facing.face(rc.robot.facing)
}

func (rc *RobotController) Move(moving Mover){
	rc.moving = moving
	rc.moving.move(rc.robot)
}
