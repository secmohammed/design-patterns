package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
)

type Address struct {
    StreetAddress, City, Country string
}

type Person struct {
    Name    string
    Address *Address
    Friends []string
}

//Better convention when having nested structs, we won't copy each struct as before
//it doesn't handle slices as well.
func (p *Person) DeepCopy() *Person {
    b := bytes.Buffer{}
    e := gob.NewEncoder(&b)
    _ = e.Encode(p)
    fmt.Println(string(b.Bytes()))
    d := gob.NewDecoder(&b)
    result := Person{}
    _ = d.Decode(&result)
    return &result
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
