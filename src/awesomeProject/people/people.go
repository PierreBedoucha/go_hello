package people

import (
	"fmt"
)

type Person struct{
	Name string
	Age int
}

func (p Person) FullName() string {
	return p.Name
}

func (p Person) format() string {
	return fmt.Sprintf("Test format() Person")
}

//Interface
func (p Person) Work() {

}