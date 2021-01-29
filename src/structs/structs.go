package main
import "fmt"

type Framework struct { name string }
type Language struct { name string }
type DB struct { name string }

type IT struct {
	l *Language
	f *Framework
	d *DB
}

func main() {
	i := IT{&Language{"Go"}, &Framework{"Nodejs"}, &DB{"Oracle"}}
	fmt.Printf("IT dept:\nLanguage: %s\nFramework: %s\nDB: %s\n", i.l, i.f, i.d)
}
