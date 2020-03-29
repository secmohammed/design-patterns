package main

import "fmt"

type Color int
type Size int

const (
    red Color = iota
    green
    blue
)
const (
    small Size = iota
    medium
    large
)

type Product struct {
    name  string
    color Color
    size  Size
}
type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
    result := make([]*Product, 0)
    for i, v := range products {
        if v.color == color {
            result = append(result, &products[i])
        }
    }
    return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
    result := make([]*Product, 0)
    for i, v := range products {
        if v.size == size {
            result = append(result, &products[i])
        }
    }
    return result
}
func (f *Filter) FilterByColorAndSize(products []Product, size Size, color Color) []*Product {
    result := make([]*Product, 0)
    for i, v := range products {
        if v.size == size && v.color == color {
            result = append(result, &products[i])
        }
    }
    return result
}

// Fixing issue using Spec pattern.
type Specification interface {
    IsSatisfied(p *Product) bool
}
type ColorSpecification struct {
    color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
    return p.color == c.color
}

type SizeSpecification struct {
    size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
    return p.size == s.size
}

type BetterFilter struct {
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
    fmt.Printf("%+v", spec)
    result := make([]*Product, 0)
    for i, v := range products {
        if spec.IsSatisfied(&v) {
            result = append(result, &products[i])
        }
    }
    return result
}

type AndSpecification struct {
    first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
    return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}
func main() {
    apple := Product{"Apple", green, small}
    tree := Product{"Tree", green, large}
    house := Product{"House", blue, large}
    products := []Product{apple, tree, house}
    fmt.Printf("Green Products (old):\n")
    f := Filter{}
    for _, v := range f.FilterByColor(products, green) {
        fmt.Printf("- %s is green \n", v.name)
    }
    fmt.Printf("green products (new way): \n")
    greenSpec := ColorSpecification{green}
    bf := BetterFilter{}
    for _, v := range bf.Filter(products, greenSpec) {
        fmt.Printf("- %s is green \n", v.name)
    }
    largeSpec := SizeSpecification{large}
    lgSpec := AndSpecification{greenSpec, largeSpec}
    fmt.Printf("Large green products: \n")
    for _, v := range bf.Filter(products, lgSpec) {
        fmt.Printf("- %s is green and large \n", v.name)

    }

}
