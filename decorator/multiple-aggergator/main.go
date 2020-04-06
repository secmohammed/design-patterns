package main

import (
    "fmt"
)

// type Bird struct {
//     Age int
// }

// func (b *Bird) Fly() {
//     if b.Age >= 10 {
//         fmt.Println("flying..")
//     }
// }

// type Lizard struct {
//     Age int
// }

// aggregator can be used this way, but that's only a primitive solution
// as if they interesect with props such as age in this scenario. It iwll blow up
//
// type Dragon struct {
//     Bird
//     Lizard
// }
// func (l *Lizard) Crawl() {
//     if l.Age < 10 {
//         fmt.Println("Crawling!")
//     }
// }
// func (d *Dragon) Age() int {
//     return d.Bird.Age
// }
// func (d *Dragon) SetAge(age int) {
//     d.Bird.Age = age
//     d.Lizard.Age = age
// }
type Aged interface {
    Age() int
    SetAge(age int)
}
type Bird struct {
    age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }
func (b *Bird) Fly() {
    if b.age >= 10 {
        fmt.Println("Flying..")
    }
}

type Lizard struct {
    age int
}

func (l *Lizard) Crawl() {
    if l.age < 10 {
        fmt.Println("Crawling...")
    }
}

type Dragon struct {
    bird   Bird
    lizard Lizard
}

func (d *Dragon) Age() int {
    return d.bird.age
}
func (d *Dragon) SetAge(age int) {
    d.lizard.SetAge(age)
    d.bird.SetAge(age)
}
func (d *Dragon) Fly() {
    d.bird.Fly()
}
func (d *Dragon) Crawl() {
    d.lizard.Crawl()
}
func (l *Lizard) Age() int {
    return l.age
}
func (l *Lizard) SetAge(age int) {
    l.age = age
}
func NewDragon() *Dragon {
    return &Dragon{Bird{}, Lizard{}}
}
func main() {
    // inconsistency. due to we have a single field and we have to set it accordingly. rather than setting by aggregator.
    //
    // d  := Dragon{}
    // d.Bird.Age = 11
    d := NewDragon()
    d.SetAge(10)
    d.Fly()
    d.Crawl()
    d.SetAge(5)
    d.Fly()
    d.Crawl()

}
