package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"

type Facer interface{
	Face(*event.Facing)
}

type FaceRight struct{}

func (fr *FaceRight) Face(facing *event.Facing){
	if facing.Direction() == event.West{
		facing.FirstDirection()
	}else{
		facing.PlusDirection()
	}
}

type FaceLeft struct{}

func (fl *FaceLeft) Face(facing *event.Facing){
	if facing.Direction() == event.North{
		facing.LastDirection()
	}else{
		facing.MinusDirection()
	}
}


