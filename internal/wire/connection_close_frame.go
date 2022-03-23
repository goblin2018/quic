package wire

import (
	"bytes"
	"quic/internal/protocol"
	"quic/quicvarint"
)

type ConnectionCloseFrame struct {
	IsApplicationError bool
	ErrorCode          uint64
	FrameType          uint64
	ReasonPhrase       string
}

func parseConnectionCloseFrame(r *bytes.Reader, _ protocol.VersionNumber) (*ConnectionCloseFrame, error) {
	typeByte, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	f := &ConnectionCloseFrame{IsApplicationError: typeByte == 0x1d}
	ec, err := quicvarint.Read(r)

	if err != nil {
		return nil, err
	}

	f.ErrorCode = ec

}
