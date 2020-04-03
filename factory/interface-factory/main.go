package main

import "fmt"

type Person interface {
    SayHello()
}
type person struct {
    name string
    age  int
}
type tiredPerson struct {
    name string
    age  int
}

func (p *person) SayHello() {
    fmt.Printf("Hi, my name is %s I am %d years old \n", p.name, p.age)

}
func (p *tiredPerson) SayHello() {
    fmt.Println("Hi, I'm too tired to talk to you.")
}

// return interface.
// we don't need to return a pointer. *Person
func NewPerson(name string, age int) Person {
    if age > 100 {
        return &tiredPerson{name, age}
    }
    return &person{name, age}
}
func main() {
    p := NewPerson("James", 32)
    // the goal of returning an interface instead of the struct itself
    // is we cannot write the data anymore
    // that means it's encapsulated in a way that's immutable.
    // also we can deal with this struct only through the methods we only have.
    p.SayHello()
    p1 := NewPerson("James", 120)
    p1.SayHello()
}
