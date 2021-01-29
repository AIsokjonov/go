package main
import (
	"fmt"
	"network"
	"structs"
)

func main() {
	fmt.Println(network.Request())

	c := structs.NewCar("BMW", "BMW X5", "Black", 73.999)
	fmt.Printf("Car:\nName: %v\nModel: %v\nColor: %v\nPrice: %v\n", c.Name, c.Model, c.Color, c.Price)


	// structs with new()
	x := new(structs.Bar)
	(*x).ThingOne = "Hello"
	(*x).ThingTwo = 1

	fmt.Println("Structs with new()")
	fmt.Printf("\nThingOne field: %v\n", x.ThingOne)
	fmt.Printf("ThingTwo field: %v\n", x.ThingTwo)

	// structs with make()
	// y := make(structs.Bar)
	// fmt.Println("\nStructs with make()")
	// fmt.Println(y)

	// reference type with make()
	z := make(structs.Foo)
	z["first"] = "Golang"
	z["second"] = "C++"

	fmt.Println("\nReference types with make()")
	fmt.Printf("First value: %v\n", z["first"])
	fmt.Printf("Second value: %v\n", z["second"])

	// w := new(structs.Foo)
	// (*w)["first"] = "Nodejs"
	// (*w)["second"] = "Java"
	// fmt.Println("Reference types with new()")
	// fmt.Printf("First value: %v\n", w["first"])
	// fmt.Printf("Second value: %v\n", w["second"])

	// anynomous struct fields
	d := new(structs.Developer)
	d.Name = "John"
	d.Position = "Backend"
	d.Framework.Name = "Nodejs"
	d.Framework.Pl = "JS"
	fmt.Println(d)
}
