package main

import "fmt"

type Driven interface {
    Drive()
}
type Car struct {
}

func (c Car) Drive() {
    fmt.Println("car is being driven.")
}
func (c *CarProxy) Drive() {
    if c.driver.Age >= 16 {
        c.car.Drive()
    } else {
        fmt.Println("Driver is too young to drive.")
    }
}

type Driver struct {
    Age int
}
type CarProxy struct {
    car    Car
    driver *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
    return &CarProxy{
        car:    Car{},
        driver: driver,
    }
}
func main() {
    car := NewCarProxy(&Driver{22})
    car.Drive()
}
