package utils

import (
	"crypto/rand"
	"encoding/binary"
)

type Rand struct {
	buf [4]byte
}

func (r *Rand) Int31() int32 {
	rand.Read(r.buf[:])
	return int32(binary.BigEndian.Uint32(r.buf[:]) & ^uint32(1<<31))
}
func (r *Rand) Int31n(n int32) int32 {
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > max {
		v = r.Int31()
	}
	return v % n
}
