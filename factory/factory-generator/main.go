package main

import "fmt"

type Employee struct {
    Name, Position string
    AnnualIncome   int
}

// functional approache.
func NewEmployeeFactory(position string, annaulIncome int) func(name string) *Employee {
    return func(name string) *Employee {
        return &Employee{name, position, annaulIncome}
    }
}

// structural approache.
type EmployeeFactory struct {
    Position     string
    AnnualIncome int
}

func NewEmployeeFactory2(position string, annaulIncome int) *EmployeeFactory {
    return &EmployeeFactory{position, annaulIncome}
}
func (f *EmployeeFactory) Create(name string) *Employee {
    return &Employee{name, f.Position, f.AnnualIncome}
}
func main() {
    developerFactory := NewEmployeeFactory("developer", 1000)
    managerFactory := NewEmployeeFactory("manager", 8000)
    developer := developerFactory("Adam")
    // In case of changing attribute.
    developer.Name = "Hani"
    manager := managerFactory("Jane")
    fmt.Println(developer, manager)
    bossFactory := NewEmployeeFactory2("CEO", 100000)
    boss := bossFactory.Create("Sami")
    fmt.Println(boss)
}
