package utils

import "sync/atomic"

type AtomicBool struct {
	v int32
}

func (a *AtomicBool) Set(value bool) {
	var n int32
	if value {
		n = 1
	}
	atomic.StoreInt32(&a.v, n)
}

func (a *AtomicBool) Get() bool {
	return atomic.LoadInt32(&a.v) != 0
}
