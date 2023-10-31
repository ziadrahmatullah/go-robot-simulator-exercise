package entity

type Mover interface{
	move(*Robot)
}

type MoveAdvance struct{}

func (ma *MoveAdvance) move(robot *Robot){
	switch robot.facing.Direction() {
	case 1:
		robot.coordinate.YCordinatePlus()
	case 2:
		robot.coordinate.XCordinatePlus()
	case 3:
		robot.coordinate.YCordinateMinus()
	case 4:
		robot.coordinate.XCordinateMinus()
	default:
	}
}