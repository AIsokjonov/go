package main
import (
	"fmt"
	"network"
	"structs"
)

func main() {
	fmt.Println(network.Request())

	var dev *structs.Developer = new(structs.Developer)
	dev.name = "James"
	dev.position = "Backend"
	fmt.Println(dev)
}
