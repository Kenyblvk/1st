package main

import "fmt"

type Person struct {
	lastname  string
	firstname string
	age       int
}

func (p Person) userdata() []Person {
	x := make([]Person, 0)
	x = append(x, p)
	return x
}
func main() {
	p := Person{
		lastname:  "Ezeilo",
		firstname: "Kenechukwu",
		age:       24,
	}
	result := p.userdata()
	fmt.Println(result)

}
