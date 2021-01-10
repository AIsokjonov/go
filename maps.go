package main
import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Printf("Map values: %T\n", m)

	// assigning int 66 to "route" key
	m["route"] = 66

	// assigning int 23 to "age" key
	m["age"] = 23

	// assigning int 52 to "population" key
	m["population"] = 52000000;

	fmt.Println("Original:", m)

	// removing elements of the map
	delete(m, "population")
	fmt.Println("After delete():", m)

	// a two-value assignment test
	i, ok := m["route"]
	fmt.Println("If value exists:", ok)
	fmt.Println("Lookup value: ", i)

	// to test for a key without retrieving the value
	// we can use _(underscore)
	_, ifExists := m["population"]
	fmt.Println("Does exist?:", ifExists)

	// length of map
	length := len(m)
	fmt.Println("Length of map:", length)

	// iterate over the contents of a map
	// use range
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	// how to initialize maps
	students := map[string]int{
		"IT": 13,
		"BBA": 19,
	}
	fmt.Println("Students map:", students)

	// exploiting zero-values
	type Node struct {
		Next	*Node
		Value	interface{}
	}
	var first *Node

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.value)
	}
}