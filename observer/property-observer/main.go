package main

import (
    "container/list"
    "fmt"
)

type Observerable struct {
    subs *list.List
}

func (o *Observerable) Subscribe(x Observer) {
    o.subs.PushBack(x)
}
func (o *Observerable) Unsubscribe(x Observer) {
    for z := o.subs.Front(); z != nil; z = z.Next() {
        if z.Value.(Observer) == x {
            o.subs.Remove(z)
        }
    }
}

func (o Observerable) Fire(data interface{}) {
    for z := o.subs.Front(); z != nil; z = z.Next() {
        z.Value.(Observer).Notify(data)
    }
}

type PropertyChange struct {
    Name  string // "Age" "Height" property to track change.
    Value interface{}
}
type Observer interface {
    Notify(data interface{})
}
type Person struct {
    Observerable
    age int
}

func (p *Person) Age() int {
    return p.age
}
func (p *Person) SetAge(age int) {
    if age == p.age {
        return
    }
    p.age = age
    p.Fire(PropertyChange{"Age", p.age})
}

type TrafficManagement struct {
    o Observerable
}

func (t *TrafficManagement) Notify(data interface{}) {
    if pc, ok := data.(PropertyChange); ok {
        if pc.Value.(int) >= 16 {
            fmt.Println("Congrats, you can drive now!")
            t.o.Unsubscribe(t)
        }
    }
}

func NewPerson(age int) *Person {
    return &Person{
        Observerable: Observerable{list.New()},
        age:          age,
    }
}

func main() {
    p := NewPerson(15)
    t := &TrafficManagement{p.Observerable}
    p.Subscribe(t)
    for i := 16; i <= 20; i++ {
        fmt.Println("Setting the age to", i)
        p.SetAge(i)
    }
}
