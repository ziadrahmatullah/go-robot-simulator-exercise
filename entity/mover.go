package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"

type Mover interface{
	move(*Robot)
}

type MoveAdvance struct{}

func (ma *MoveAdvance) move(robot *Robot){
	switch robot.facing.Direction() {
	case event.North:
		robot.coordinate.YCordinatePlus()
	case event.East:
		robot.coordinate.XCordinatePlus()
	case event.South:
		robot.coordinate.YCordinateMinus()
	case event.West:
		robot.coordinate.XCordinateMinus()
	default:
	}
}

func IsMoveCommand(cmd string) bool{
	return cmd == "A"
}