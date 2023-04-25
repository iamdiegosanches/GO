package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

type Account struct {
	balance int
	lock sync.Mutex // Only allow one customer to access the account
}

func (a *Account) GetBallance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enouth money in account")
	} else {
		fmt.Printf("%d withdrawn : Balacane: %d\n", v, a.balance)
		a.balance -= v
	}
}

func main() {
	var acct Account
	acct.balance = 100
	pl("Balance:", acct.GetBallance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2* time.Second)

}
