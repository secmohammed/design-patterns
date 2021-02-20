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

type Observer interface {
    Notify(data interface{})
}
type Person struct {
    Observerable
    Name string
}

func NewPerson(name string) *Person {
    return &Person{
        Observerable: Observerable{new(list.List)},
        Name:         name,
    }
}
func (p *Person) CatchACold() {
    p.Fire(p.Name)
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
    fmt.Printf("A Doctor has been called for %s", data.(string))
}
func main() {
    p := NewPerson("Renaldo")
    ds := &DoctorService{}
    p.Subscribe(ds)
    p.CatchACold()
    p.Unsubscribe(ds)
    p.CatchACold()

}
