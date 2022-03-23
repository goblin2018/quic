package utils

import (
	"math"
	"time"
)

type Timer struct {
	t        *time.Timer
	read     bool
	deadline time.Time
}

func NewTimer() *Timer {
	return &Timer{t: time.NewTimer(time.Duration(math.MaxInt64))}
}

func (t *Timer) Chan() <-chan time.Time {
	return t.t.C
}

func (t *Timer) Reset(deadline time.Time) {
	if deadline.Equal(t.deadline) && !t.read {
		return
	}

	if !t.t.Stop() && !t.read {
		<-t.t.C
	}
	if !deadline.IsZero() {
		t.t.Reset(time.Until(deadline))
	}
	t.read = false
	t.deadline = deadline
}

func (t *Timer) SetRead() {
	t.read = true
}
func (t *Timer) Stop() {
	t.t.Stop()
}
