package utils

import (
	"math"
	"time"

	"golang.org/x/exp/constraints"
)

const InfDuration = time.Duration(math.MaxInt64)

func Max[T constraints.Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Abs[T constraints.Signed](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

func MinTime(a, b time.Time) time.Time {
	if a.After(b) {
		return b
	}
	return a
}
func MinNonZeroTime(a, b time.Time) time.Time {
	if a.IsZero() {
		return b
	}
	if b.IsZero() {
		return a
	}
	return MinTime(a, b)
}

// MaxTime returns the later time
func MaxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
