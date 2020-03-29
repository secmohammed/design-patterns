package main

type Document struct {
}
type Printer interface {
    Print(d Document)
}
type Scanner interface {
    Scan(d Document)
}
type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {

}
func (m MultiFunctionPrinter) Fax(d Document) {

}
func (m MultiFunctionPrinter) Scan(d Document) {

}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct {
}

func (p Photocopier) Print(d Document) {

}
func (p Photocopier) Scan(d Document) {

}

//Composition of interfaces.
type MultiFunctionDevice interface {
    Printer
    Scanner
}

// Decoarator Design Pattern.
type MultiFunctionMachine struct {
    printer Printer
    scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
    m.printer.Print(d)
}
func (m MultiFunctionMachine) Scan(d Document) {
    m.scanner.Scan(d)
}
func main() {
    document := Document{}
    multifunction := MultiFunctionMachine{
        printer: MyPrinter{},
        scanner: Photocopier{},
    }
    multifunction.Print(document)
}
