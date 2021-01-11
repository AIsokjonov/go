packgge main
import "fmt"

type developer struct {
	position	string
	skill		string
}

var m map[string]developer{
	"James": developer {
		"backend", "Go"
	},
	"Bob": developer {
		"mobile", "Java"
	}
}

func main() {
	fmt.Println(m["James"])
}
