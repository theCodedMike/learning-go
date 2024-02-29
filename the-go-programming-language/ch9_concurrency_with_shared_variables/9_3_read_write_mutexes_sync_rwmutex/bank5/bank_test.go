package bank_test

// 在终端执行：
//
//	go test -v ./ch9_concurrency_with_shared_variables/9_3_read_write_mutexes_sync_rwmutex/bank5

import (
	bank "gopl.io/ch9_concurrency_with_shared_variables/9_3_read_write_mutexes_sync_rwmutex/bank5"
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()
	if got, want := bank.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Withdraw [1..1000] concurrently.
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			if !bank.Withdraw(amount) {
				t.Errorf("Failed to withdraw: %d", amount)
			}
			n.Done()
		}(i)
	}
	n.Wait()
	if got := bank.Balance(); got != 0 {
		t.Errorf("Balance = %d, want %d", got, 0)
	}
}
