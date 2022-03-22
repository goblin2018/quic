package wire

import (
	"bytes"
	"errors"
	"quic/internal/protocol"
	"time"
)

var errInvalidAckRanges = errors.New("AckFrame: ACK frame contains invalid ACK ranges")

type AckFrame struct {
	AckRanges         []AckRange
	DelayTime         time.Duration
	ECT0, ECT1, ECNCE uint64
}

func parseFrame(r *bytes.Reader, ackDelayExponent uint8, _ protocol.VersionNumber) (*AckFrame, error) {
	typeByte, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	ecn := typeByte&0x1 > 0
	frame := &AckFrame{}

}
