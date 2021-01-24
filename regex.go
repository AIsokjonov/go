package main
import (
	"strconv"
	"fmt"
	"regexp"
)

func main() {
	searchStr := "Hello, My name is Abdurauf and I'm a 24.0` y.o software developer from Uzbekistan. Currently I live and work in Seoul, South Korea"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v * 2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchStr)); ok {
		fmt.Println("Match found")
	}
	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(searchStr, "##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchStr, f)
	fmt.Println(str2)
}