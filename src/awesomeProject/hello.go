package main

import (
	"errors"
	"fmt"

	"~/go/src/awesomeProject/people"
)

type Thing string

type Worker interface {
	Work()
}

func main() {
	fmt.Printf("hello, world\n")

	name := "Pierre"
	//var foo string
	//thing := Thing("foo")

	fmt.Printf("hello, %s\n", name)

	p := people.Person{
		Name: "PierreB",
		Age: 30,
	}
	fmt.Printf("hello, %s\n", p.Name)

	//calling the function
	greeting := formatPerson(p)
	fmt.Println(greeting)

	//calling method
	//fmt.Println(p.FullName())

	//create a slice, args: type, initial size, capacity
	slice := make([]person, 0, 10)
	fmt.Printf("length=%d\n", len(slice))

	//add data to the slice
	slice = append(slice, person{name:"Test", age: 3})
	//slice iteration
	for _, p1:= range slice{
		//slice[i].format() or use _ to void it to avaid error
		p1.FullName()
	}

	//Maps
	aMap := make(map[string]person)
	aMap["Foo"] = person{}
	fmt.Println("$v", aMap["Foo"])
	//for wrong key, the zero-value is returned
	bar, ok := aMap["bar"]
	if !ok{
		panic("oh no!")
	}

	formattedName, err := formatPerson(p)
	if err != nil {
		panic(err)
	}

	//Pointers
	point := &p

}

//Interface
func doSomething(p Worker) {
	p.Work()
}

//pointers and function
func thing (p *people.Person) {
	p.Name = "point!"
}

//functions!
func formatPerson(p person) (string, error) {
	// first varname, then type and then return type

	if p.name == "" {
		return "", errors.New("No Name!")
	}

	return fmt.Sprintf("hello, %s\n", p.Name), nil
}

// Method
//func (p person) FullName() string {
//	return p.name
//}

//Concurrency
//sync.WaitGroup, to wait for one unit. need to pass a pointer to the func and call .done()
//channel:
//chan, read (will block it), and write to it (same, block)
// use <-c to read and c<-true to write (true here)

//HTTP
//server and client out of the box
//import net/http package






