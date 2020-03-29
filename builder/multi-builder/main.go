package main

import "fmt"

type Person struct {
    StreetAddress, Postcode, City string
    // job
    CompanyName, Position string
    AnnualIncome          int
}

type PersonBuilder struct {
    person *Person
}
type PersonAddressBuilder struct {
    PersonBuilder
}
type PersonJobBuilder struct {
    PersonBuilder
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
    return &PersonAddressBuilder{*b}
}
func (b *PersonBuilder) Works() *PersonJobBuilder {
    return &PersonJobBuilder{*b}
}
func NewPersonBuilder() *PersonBuilder {
    return &PersonBuilder{&Person{}}
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
    pjb.person.Position = position
    return pjb
}
func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
    pjb.person.AnnualIncome = annualIncome
    return pjb
}
func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
    pjb.person.CompanyName = companyName
    return pjb
}
func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
    b.person.StreetAddress = streetAddress
    return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
    b.person.City = city
    return b
}
func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
    b.person.Postcode = postcode
    return b
}
func (b *PersonBuilder) Build() *Person {
    return b.person
}
func main() {
    pb := NewPersonBuilder()
    pb.Lives().At("123 London Road").In("London").WithPostcode("SW12BC").Works().At("Fabrikam").AsA("Programmer").Earning(123000)
    person := pb.Build()
    fmt.Println(person)
}
