package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"

type Facer interface{
	face(*event.Facing)
}

type FaceRight struct{}

func (fr *FaceRight) face(facing *event.Facing){
	if facing.Direction() == event.West{
		facing.FirstDirection()
	}else{
		facing.PlusDirection()
	}
}

type FaceLeft struct{}

func (fl *FaceLeft) face(facing *event.Facing){
	if facing.Direction() == event.North{
		facing.LastDirection()
	}else{
		facing.MinusDirection()
	}
}

func IsFacingCommand(cmd string) bool{
	return cmd == "R" || cmd == "L"
}
