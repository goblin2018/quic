package protocol

import (
	"fmt"
	"time"
)

type PacketType uint8

const (
	PacketTypeInitial PacketType = 1 + iota
	PacketTypeRetry
	PacketTypeHandshake
	PacketType0RTT
)

func (t PacketType) String() string {
	switch t {
	case PacketTypeInitial:
		return "Initial"
	case PacketTypeRetry:
		return "Retry"
	case PacketTypeHandshake:
		return "Handshake"
	case PacketType0RTT:
		return "0-RTT Protected"
	default:
		return fmt.Sprintf("unknown packet type: %d", t)
	}
}

type ECN uint8

const (
	ECNNon ECN = iota
	ECT1
	ECT0
	ECNCE
)

type ByteCount int64

const MaxByteCount = ByteCount(1<<62 - 1)

const InvalidByteCount ByteCount = -1

type StatelessResetToken [16]byte

// 以太网 1500
// UDP 8
// IPv6 40
const MaxPacketBufferSize ByteCount = 1452
const MinInitialPacketSize = 1200

const MinUnknownVersionPacketSize = MinInitialPacketSize
const MinStatelessResetSize = 1 /* first byte */ + 20 /* max. conn ID length */ + 4 /* max. packet number length */ + 1 /* min. payload length */ + 16 /* token */
const MinConnectionIDLenInitial = 8
const DefaultAckDelayExponent = 3
const MaxAckDelayExponent = 20
const DefaultMaxAckDelay = 25 * time.Millisecond
const MaxMaxAckDelay = (1<<14 - 1) * time.Millisecond
const MaxConnIDLen = 20
const InvalidPacketLimitAES = 1 << 52
const InvalidPacketLimitChaCha = 1 << 36
