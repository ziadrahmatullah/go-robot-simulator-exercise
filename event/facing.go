package event

type Facing struct {
	direction int
}

func NewFacing(direction string) *Facing{
	switch direction{
	case "N":
		return &Facing{direction: 1}
	case "E":
		return &Facing{direction: 2}
	case "S":
		return &Facing{direction: 3}
	case "W":
		return &Facing{direction: 4}
	default:
		return nil
	}
}

func (f *Facing) Direction() int {
	return f.direction
}

func (f *Facing) PlusDirection() {
	f.direction++
}

func (f *Facing) MinusDirection() {
	f.direction--
}

func (f *Facing) FirstDirection() {
	f.direction = 1
}

func (f *Facing) LastDirection() {
	f.direction = 4
}
