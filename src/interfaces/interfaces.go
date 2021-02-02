package main
import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}

func main() {
	c := new(Customer)
	c.Name = "Borack Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	fmt.Println(c.Log())

	// shorter
	c2 := &Customer{"Donal Trump", &Log{"Make America great Again!"}}
	fmt.Println(c2.log)
	c2.Log().Add("2 - After me, the world will be a better place!")
	fmt.Println(c2.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
