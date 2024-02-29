package bank_test

// 在终端执行：
//
//	go test -v ./ch9_concurrency_with_shared_variables/9_2_mutual_exclusion_sync_mutex/bank2

import (
	bank "gopl.io/ch9_concurrency_with_shared_variables/9_2_mutual_exclusion_sync_mutex/bank2"
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
}
