package main
import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}

func main() {
	c := &Customer{"Donald Trump", Log{"Make America Great Again!"}}
	fmt.Println(c.msg)
	c.Log.Add("by Donald Trump")
	fmt.Println(c.msg)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}
