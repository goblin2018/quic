package utils

import "quic/internal/protocol"

type ByteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}
