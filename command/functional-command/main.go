package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
    Balance int
}

func Deposit(b *BankAccount, amount int) {
    b.Balance += amount
    fmt.Println("Deposited", amount, "\b balance is now", b.Balance)
}
func Withdraw(b *BankAccount, amount int) bool {
    if b.Balance-amount >= overdraftLimit {
        b.Balance -= amount
        fmt.Println("Withdrew", amount, "\b Balance is now", b.Balance)
        return true
    }
    return false
}

func main() {
    ba := &BankAccount{0}
    var commands []func()
    commands = append(commands, func() {
        Deposit(ba, 100)
    })
    commands = append(commands, func() {
        Withdraw(ba, 25)
    })
    for _, cmd := range commands {
        cmd()
    }
    fmt.Println(ba)

}
