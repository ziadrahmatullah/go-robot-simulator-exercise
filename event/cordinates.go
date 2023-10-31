package event

type Cordinates struct{
	xAxis, yAxis int
}

func NewCordinate(xAxis, yAxis int)*Cordinates{
	return &Cordinates{xAxis: xAxis, yAxis: yAxis}
}

func (c *Cordinates) XAxis()int{
	return c.xAxis
}

func (c *Cordinates) YAxis() int{
	return c.yAxis
}

func (c *Cordinates) XCordinatePlus(){
	c.xAxis++
}

func (c *Cordinates) YCordinatePlus(){
	c.yAxis++
}

func (c *Cordinates) XCordinateMinus(){
	c.xAxis--
}

func (c *Cordinates) YCordinateMinus(){
	c.yAxis--
}



