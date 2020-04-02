package main

type Person struct {
    Nme      string
    Age      int
    EyeCount int
}

func NewPerson(name string, age int) *Person {
    // sometimes we need to process rules while instantiating the struct.
    // if age < 16  {
    //     // ..
    // }
    return &Person{name, age, 2}
}
func main() {
    p := NewPerson("John", 22)
    // in case of the need to overwrite.
    p.EyeCount = 1
}
