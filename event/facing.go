package event

type Facing struct {
	direction int
}

const(
	North = iota +1
	East
	South
	West
)

func NewFacing(direction string) *Facing{
	switch direction{
	case "N":
		return &Facing{direction: North}
	case "E":
		return &Facing{direction: East}
	case "S":
		return &Facing{direction: South}
	case "W":
		return &Facing{direction: West}
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


