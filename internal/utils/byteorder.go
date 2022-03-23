package utils

import (
	"bytes"
	"io"
)

type ByteOrder interface {
	ReadUint32(io.ByteReader) (uint32, error)
	ReadUint24(io.ByteReader) (uint32, error)
	ReadUint16(io.ByteReader) (uint16, error)

	WriteUint32(*bytes.Buffer, uint32)
	WriteUint24(*bytes.Buffer, uint32)
	WriteUint16(*bytes.Buffer, uint16)
}
