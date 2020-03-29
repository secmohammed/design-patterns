package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

var entryCount = 0

type Journal struct {
    entries []string
}

func (j *Journal) AddEntry(text string) int {
    entryCount++
    entry := fmt.Sprintf(
        "%d: %s",
        entryCount,
        text,
    )
    j.entries = append(j.entries, entry)
    return entryCount
}
func (j *Journal) String() string {
    return strings.Join(j.entries, "\n")
}

// breaks s
func (j *Journal) Save(filename string) {
    _ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

// better convention
// at another package.
var LineSpearator = "\n"

func SaveToFile(j *Journal, filename string) {
    _ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

type Persistence struct {
    lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
    _ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

// breaks s
func (j *Journal) Load(filename string) {

}

func main() {
    j := Journal{}
    j.AddEntry("I cried today")
    j.AddEntry("I ate a bug")
    fmt.Println(j.String())
    SaveToFile(&j, "journal.txt")
    p := Persistence{"\r\n"}
    p.SaveToFile(&j, "journal.txt")
}
