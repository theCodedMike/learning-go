package memo_test

import (
	memo "gopl.io/ch9_concurrency_with_shared_variables/9_7_example_concurrent_non_blocking_cache/memo1"
	"gopl.io/ch9_concurrency_with_shared_variables/9_7_example_concurrent_non_blocking_cache/memotest"
	"testing"
)

// 在终端执行：
//
//  go test -run=TestConcurrent -race -v ./ch9_concurrency_with_shared_variables/9_7_example_concurrent_non_blocking_cache/memo2

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
