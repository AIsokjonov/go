package methods

type Camera struct { }

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct { }

func (p *Phone) Call() string {
	return "Ring Ring"
}

type Smartphone struct {
	Camera
	Phone
}
