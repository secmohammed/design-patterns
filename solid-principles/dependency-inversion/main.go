package main

import "fmt"

type Relationship int

const (
    Parent Relationship = iota
    Child
    Sibling
)

type Person struct {
    name string
}
type Info struct {
    from         *Person
    relationship Relationship
    to           *Person
}

// Low Level module.
type Relationships struct {
    relations []Info
}

// Low Level module.
type RelationshipBrowser interface {
    FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
    // since we are at a low level, we can use the other low level directly as this doesn't violate the principle.
    result := make([]*Person, 0)
    for i, v := range r.relations {
        if v.relationship == Parent && v.from.name == name {
            result = append(result, r.relations[i].to)
        }
    }
    return result
}
func (r *Relationships) AddParentAndChild(parent, child *Person) {
    r.relations = append(r.relations, Info{
        parent, Parent, child,
    })
}

// High level module.
type Research struct {
    //Breaks DIP. as we depend on low level module.
    // relationships Relationships
    // doesn't break the DIP as we depend on concrete.
    browser RelationshipBrowser
}

// Old
// func (r *Research) Investigate() {
//     relations := r.relationships.relations
//     for _, rel := range relations {
//         if rel.from.name == "John" && rel.relationship == Parent {
//             fmt.Println("John has a child called ", rel.to.name)
//         }
//     }
// }
func (r *Research) Investigate() {
    for _, p := range r.browser.FindAllChildrenOf("Jon") {
        fmt.Println("John has a child called", p.name)
    }
}

func main() {
    parent := Person{"John"}
    child := Person{"Chris"}
    child2 := Person{"Matt"}
    relationships := Relationships{}
    relationships.AddParentAndChild(&parent, &child)
    relationships.AddParentAndChild(&parent, &child2)
    r := Research{&relationships}
    r.Investigate()
}
