package structs

type Car struct {
	Name string
	Model string
	Color string
	Price float64
}

func NewCar(Name string, Model string, Color string, Price float64) *Car {
	c := new(Car)
	return c
}

type Foo map[string]string

type Bar struct {
	ThingOne string
	ThingTwo int
}

type Framework struct {
	Name string
	Pl string
}

// struct anonymous fields
type Developer struct {
	Name string
	Position string

	// string // anynomous field
	Framework // anynomous field
}
