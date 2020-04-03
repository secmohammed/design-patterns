package main

import "fmt"

type Address struct {
    StreetAddress, City, Country string
}
type Person struct {
    Name    string
    Address *Address
}

func main() {
    john := Person{
        "John",
        &Address{"123 London rd", "London", "UK"},
    }
    // this causes an issue, due to pointing.
    // we copied the pointer of the address as well, which will affect both john and jane.
    // jane := john
    // jane.Name = "Jane"
    // jane.Address.StreetAddress = "321 Baker St"
    //    fmt.Println(john, john.Address)
    //    fmt.Println(jane, jane.Address)

    jane := john
    // create a new pointer address to the address with the content of the prvious address.
    jane.Address = &Address{
        john.Address.StreetAddress,
        john.Address.City,
        john.Address.Country,
    }
    jane.Name = "Jane"
    jane.Address.StreetAddress = "321 Baker St"

    fmt.Println(john, john.Address)
    fmt.Println(jane, jane.Address)
}
