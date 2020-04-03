package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
)

type Address struct {
    Suite               int
    StreetAddress, City string
}

type Employee struct {
    Name   string
    Office Address
}

func (p *Employee) DeepCopy() *Employee {
    b := bytes.Buffer{}
    e := gob.NewEncoder(&b)
    _ = e.Encode(p)
    fmt.Println(string(b.Bytes()))
    d := gob.NewDecoder(&b)
    result := Employee{}
    _ = d.Decode(&result)
    return &result
}

var mainOffice = Employee{
    "", Address{0, "123 East Dr", "London"},
}
var auxOffice = Employee{
    "", Address{0, "66 West Dr", "London"},
}

func NewEmployee(proto *Employee, name string, suite int) *Employee {
    result := proto.DeepCopy()
    result.Name = name
    result.Office.Suite = suite
    return result
}
func NewMainOfficeEmployee(name string, suite int) *Employee {
    return NewEmployee(&mainOffice, name, suite)
}
func NewAuxOfficeEmployee(name string, suite int) *Employee {
    return NewEmployee(&auxOffice, name, suite)
}
func main() {
    john := NewMainOfficeEmployee("john", 100)
    jane := NewAuxOfficeEmployee("jane", 100)
    fmt.Println(john, jane)
}
