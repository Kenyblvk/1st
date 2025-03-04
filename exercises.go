package main

import "fmt"

type Car struct {
	Make    string
	Model   string
	Year    int
	Running bool
}

func (c *Car) Start() {
	c.Running = true
}
func (c *Car) Stop() {
	c.Running = false
}
func (c Car) IsRunning() bool {
	return c.Running
}
func main() {
	c := Car{Make: "Kia",
		Model:   "Kn",
		Year:    1834,
		Running: true}

	c.Start()
	fmt.Println(c.IsRunning())
	c.Stop()
	fmt.Println(c.IsRunning())

}
