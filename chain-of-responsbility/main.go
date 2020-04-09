package main

import "fmt"

type Creature struct {
    Name            string
    Attack, Defense int
}

func (c *Creature) String() string {
    return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack, defense int) *Creature {
    return &Creature{name, attack, defense}
}

type IncreasedDefenseModifier struct {
    CreatureModifier
}

func NewIncreasedDefenseModifier(creature *Creature) *IncreasedDefenseModifier {
    return &IncreasedDefenseModifier{CreatureModifier{creature: creature}}
}
func (i *IncreasedDefenseModifier) Handle() {
    if i.creature.Attack <= 2 {
        fmt.Println("Increasing ", i.creature.Name, "\b's defense")
        i.creature.Defense++
    }
    i.CreatureModifier.Handle()
}

type Modifier interface {
    Add(m Modifier)
    Handle()
}
type CreatureModifier struct {
    creature *Creature
    next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
    return &CreatureModifier{creature: creature}
}

func (c *CreatureModifier) Add(m Modifier) {
    if c.next != nil {
        c.next.Add(m)
        return
    }
    c.next = m
}

func (c *CreatureModifier) Handle() {
    if c.next != nil {
        c.next.Handle()
    }
}

type DoubleAttackModifier struct {
    CreatureModifier
}

func NewDoubleAttackModifier(creature *Creature) *DoubleAttackModifier {
    return &DoubleAttackModifier{CreatureModifier{creature: creature}}
}
func (d *DoubleAttackModifier) Handle() {
    fmt.Println("Doubling", d.creature.Name, "\b's attack")
    d.creature.Attack *= 2
    d.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
    CreatureModifier
}

func NewNoBonusesModifier(creature *Creature) *NoBonusesModifier {
    return &NoBonusesModifier{CreatureModifier{creature: creature}}
}
func (n *NoBonusesModifier) Handle() {

}

func main() {
    goblin := NewCreature("Goblin", 1, 1)
    fmt.Println(goblin.String())
    root := NewCreatureModifier(goblin)
    root.Add(NewNoBonusesModifier(goblin))
    root.Add(NewDoubleAttackModifier(goblin))
    root.Add(NewIncreasedDefenseModifier(goblin))
    root.Add(NewDoubleAttackModifier(goblin))
    root.Handle()
    fmt.Println(goblin.String())
}
