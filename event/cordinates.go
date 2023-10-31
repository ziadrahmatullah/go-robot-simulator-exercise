package event

type Cordinates struct{
	xAxis, yAxis int
}

func NewCordinate(xAxis, yAxis int)*Cordinates{
	return &Cordinates{xAxis: xAxis, yAxis: yAxis}
}

func (c *Cordinates) XCordinatePlus(){
	c.xAxis++
}

func (c *Cordinates) YCordinatePlus(){
	c.xAxis++
}

func (c *Cordinates) XCordinateMinus(){
	c.xAxis--
}

func (c *Cordinates) YCordinateMinus(){
	c.xAxis--
}



