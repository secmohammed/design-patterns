package main

import "fmt"

type Address struct {
    StreetAddress, City, Country string
}

type Person struct {
    Name    string
    Address *Address
    Friends []string
}

func (p *Person) DeepCopy() *Person {
    q := *p
    q.Address = p.Address.DeepCopy()
    copy(q.Friends, p.Friends)
    return &q
}
func (a *Address) DeepCopy() *Address {
    return &Address{
        a.StreetAddress,
        a.City,
        a.Country,
    }
}
func main() {
    john := Person{
        "John",
        &Address{
            "123 London Rd", "London", "UK",
        },
        []string{"Chris", "Matt"},
    }
    jane := john.DeepCopy()
    jane.Name = "Jane"
    jane.Address.StreetAddress = "321 Baker st"
    fmt.Println(john, john.Address)
    fmt.Println(jane, jane.Address)
}
