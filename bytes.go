package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	fmt.Printf("Buffer: %v\n", &buf)

	buf.WriteString("SBV")
	fmt.Printf("Buffer: %v\n", &buf)

	buf.WriteString("ABV")
	fmt.Printf("Buffer: %v\n", &buf)
	fmt.Printf("Altogether: %v\n", buf.String())
}
