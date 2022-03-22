package protocol

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
)

type ConnectionID []byte

const maxConnectionIDLen = 20

func GenerateConnectionID(len int) (ConnectionID, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return ConnectionID(b), nil
}

func GenerateConnectionIDForInitial() (ConnectionID, error) {
	r := make([]byte, 1)
	if _, err := rand.Read(r); err != nil {
		return nil, err
	}
	len := MinConnectionIDLenInitial + int(r[0]%(maxConnectionIDLen-MinConnectionIDLenInitial+1))
	return GenerateConnectionID(len)
}

func ReadConnectionID(r io.Reader, len int) (ConnectionID, error) {
	if len == 0 {
		return nil, nil
	}
	c := make(ConnectionID, len)
	_, err := io.ReadFull(r, c)
	if err == io.ErrUnexpectedEOF {
		return nil, io.EOF
	}
	return c, err
}

func (c ConnectionID) Equal(other ConnectionID) bool {
	return bytes.Equal(c, other)
}

func (c ConnectionID) Len() int {
	return len(c)
}

func (c ConnectionID) Bytes() []byte {
	return []byte(c)
}
func (c ConnectionID) String() string {
	if c.Len() == 0 {
		return "(empty)"
	}
	return fmt.Sprintf("%x", c.Bytes())
}
