package bank_test

// 在终端执行：
//
//	go test -v ./ch9_concurrency_with_shared_variables/9_1_race_conditions/bank1

import (
	"fmt"
	"testing"
)

import (
	"gopl.io/ch9_concurrency_with_shared_variables/9_1_race_conditions/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance()) // 300
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Club
	go func() {
		ok := bank.Withdraw(100)
		fmt.Println("=", bank.Balance(), ", ok=", ok) // 200 , ok= true
		done <- struct{}{}
	}()
	<-done
	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Dick
	go func() {
		ok := bank.Withdraw(300)
		fmt.Println("=", bank.Balance(), ", ok=", ok) // 200 , ok= false
		done <- struct{}{}
	}()
	<-done
	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
