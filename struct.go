package main 
import "fmt"
type Person struct {
  lastname string 
  firstname string
  age int
  }
func (p Person ) userdata () (string, string , int ){
return p.lastname, p.firstname, p.age 
  }
func main () {
  p := Person {
    "lastname" :  "Ezeilo",
    "firstname" : "Kenechukwu",
    "age" : 24,
    }
fmt.Println(username(p))
}
    
