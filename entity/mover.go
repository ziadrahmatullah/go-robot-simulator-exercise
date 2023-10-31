package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"

type Mover interface{
	Move(*Robot) Robot
}

type MoveAdvance struct{}

func (ma *MoveAdvance) Move(robot *Robot) Robot{
	switch robot.facing.Direction() {
	case event.North:
		robot.coordinate.YCordinatePlus()
		return *robot
	case event.East:
		robot.coordinate.XCordinatePlus()
		return *robot
	case event.South:
		robot.coordinate.YCordinateMinus()
		return *robot
	default:
		robot.coordinate.XCordinateMinus()
		return *robot
	}
}