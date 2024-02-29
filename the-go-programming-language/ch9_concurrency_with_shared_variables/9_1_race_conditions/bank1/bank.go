// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan int)
var monitor = make(chan bool)

func Deposit(amount int) {
	deposits <- amount
}
func Balance() int {
	return <-balances
}
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-monitor
}
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if amount > balance {
				monitor <- false
			} else {
				balance -= amount
				monitor <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
