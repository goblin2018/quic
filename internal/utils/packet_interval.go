package utils

import "quic/internal/protocol"

type PacketInterval struct {
	Start protocol.PacketNumber
	End   protocol.PacketNumber
}
