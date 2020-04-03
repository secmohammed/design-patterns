package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "sync"
)

// better to be instantiated via factory.
type singletonDatabase struct {
    capitals map[string]int
}

func readData(path string) (map[string]int, error) {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)

    file, err := os.Open(exPath + path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    result := map[string]int{}

    for scanner.Scan() {
        k := scanner.Text()
        scanner.Scan()
        v, _ := strconv.Atoi(scanner.Text())
        result[k] = v
    }

    return result, nil
}

func (db *singletonDatabase) GetPopulation(name string) int {
    return db.capitals[name]
}

// sync.Once init() -- thread safety
// laziness can't be guaranteed when using init, but can be guaranteed using once.
// Package init() functions are guaranteed by the spec to be called
// only once and all called from a single thread
// (not to say they couldn't start goroutines, but they're thread safe unless you make them multi-threaded).
//The reason you'd use sync.Once is
//if you want to control if and when some code is executed.
//A package init() function will be called at application start,period.
// sync.Once allows you to do things like lazy initialization,
// for example creating a resource the first time it is requested
//  (but only once, in case multiple "first" requests come in at the same time)
// rather than at application start; or to only initialize a resource if it is actually going to be needed.
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
    once.Do(func() {
        caps, e := readData("./capitals.txt")
        fmt.Println(caps)
        db := singletonDatabase{caps}
        if e == nil {
            db.capitals = caps
        }
        instance = &db
    })
    return instance
}
func main() {
    db := GetSingletonDatabase()
    pop := db.GetPopulation("Seoul")
    fmt.Println("Pop of Seoul = ", pop)
}
